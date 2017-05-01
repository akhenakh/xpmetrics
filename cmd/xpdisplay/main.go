package main

import (
	"log"

	"github.com/akhenakh/xpmetrics"
)

func main() {
	l, err := xpmetrics.NewListener("127.0.0.1:49012")
	if err != nil {
		log.Fatal(err)
	}
	l.Debug = true

	l.Start()
}
