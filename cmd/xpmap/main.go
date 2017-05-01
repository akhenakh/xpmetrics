package main

import (
	"log"
	"net/http"

	"github.com/akhenakh/xpmetrics"
	"github.com/janberktold/sse"
)

// Quick & dirty moving map

// Position a data to be sent as JSON via SSE
type Position struct {
	Latitude  float32
	Longitude float32
	Altitude  float32
}

func main() {
	l, err := xpmetrics.NewListener(":49012")
	if err != nil {
		log.Fatal(err)
	}

	pChan := make(chan (*Position))

	l.Notify = func(msg xpmetrics.XPMsg, fvals [8]float32) {
		if msg == xpmetrics.LatLonAltMsg {
			p := &Position{
				Latitude:  fvals[0],
				Longitude: fvals[1],
				Altitude:  fvals[2],
			}

			select {
			case pChan <- p:
			default:
			}
		}
	}

	go l.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		conn, err := sse.Upgrade(w, r)
		if err != nil {
			return
		}
		for {
			p := <-pChan
			conn.WriteJson(p)
		}
	})
	mux.Handle("/map/", http.StripPrefix("/map/", http.FileServer(http.Dir("./static"))))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
