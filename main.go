package main

import (
	"flag"
	"log"

	"github.com/ethereum/go-ethereum/p2p"
)

func main() {
	var (
		port       = flag.Int("port", 30301, "Server port")
		connectTo  = flag.String("connect-to", "", "Enode url to connect to")
		serverName string
		err        error
		srv        *p2p.Server
	)

	flag.Parse()

	if *connectTo == "" {
		serverName = "server-1"
	} else {
		serverName = "server-2"
	}

	srv, err = newServer(serverName, *port)
	if err != nil {
		log.Printf("Failed to create new server with err: %v", err)
		return
	}

	if *connectTo != "" {
		err := connectToPeer(srv, *connectTo)
		if err != nil {
			log.Printf("Failed to connect to peer with err: %v", err)
			return
		}

		communicated := make(chan bool)
		go subscribeToEvents(srv, communicated)

		// Sent and received message, stopping the server since work is done
		<-communicated
		srv.Stop()

	} else {
		log.Println(srv.NodeInfo().Enode)

		// Running forever since nodes can reconnect
		select {}
	}

}
