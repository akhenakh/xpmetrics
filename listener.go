package xpmetrics

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"

	"github.com/pkg/errors"
)

// Listener listen for incoming X-Plane packets
type Listener struct {
	Debug bool
	Data  *XPData
	conn  *net.UDPConn
	quit  chan bool
}

// NewListener create a new Listener
func NewListener(addr string) (*Listener, error) {
	saddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:49012")
	if err != nil {
		return nil, errors.Wrap(err, "creating a new listener")
	}
	conn, err := net.ListenUDP("udp", saddr)
	if err != nil {
		return nil, errors.Wrap(err, "creating a new listener")
	}

	return &Listener{
		conn: conn,
		Data: NewXPData(),
		quit: make(chan bool),
	}, nil
}

// Start processing of new messages
func (l *Listener) Start() {
	buf := make([]byte, 65507)

	for {
		select {
		case <-l.quit:
			return
		default:
			n, _, err := l.conn.ReadFromUDP(buf)
			if err != nil {
				log.Println("can't read dataframe", err)
				return
			}

			// 5 prolog, 36 per message
			if (n-5)%36 != 0 {
				if l.Debug {
					log.Println("invalid dataframe", buf)
				}
				continue
			}

			//  4 bytes of message indicates the index number
			if bytes.Compare(buf[:4], []byte("DATA")) != 0 {
				if l.Debug {
					log.Println("invalid dataframe")
				}
				continue
			}

			for i := 0; i < (n-5)/36; i++ {
				didx, fvals, err := l.ProcessMsg(buf[5+i*36 : 5+i*36+36])
				if err != nil {
					log.Println("error processing msg", err)
				}
				if l.Debug {
					var msgType XPMsg
					msgType = XPMsg(didx)
					log.Println("line", msgType, didx, fvals)
				}

				l.Data.Insert(didx, fvals)
			}
		}
	}
}

// ProcessMsg read byte format from X-Plane and transform it to readable values
func (l *Listener) ProcessMsg(b []byte) (uint8, [8]float32, error) {
	var didx uint8
	var fvals [8]float32

	if len(b) != 36 {
		return didx, fvals, errors.New("invalid msg")
	}
	br := bytes.NewReader(b)
	err := binary.Read(br, binary.BigEndian, &didx)
	if err != nil {
		return didx, fvals, errors.Wrap(err, "invalid msg can't read idx")
	}

	_, err = br.Seek(3, 1)
	if err != nil {
		return didx, fvals, errors.Wrap(err, "invalid msg can't seek into buf")
	}

	for j := 0; j < 8; j++ {
		err = binary.Read(br, binary.LittleEndian, &fvals[j])
		if err != nil {
			return didx, fvals, errors.Wrapf(err, "can't read float val %d", j)
		}
	}

	return didx, fvals, nil
}

// Stop the listener
func (l *Listener) Stop() error {
	l.quit <- true
	return l.conn.Close()
}
