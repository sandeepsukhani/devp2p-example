package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

// Creating new p2p server
func newServer(name string, port int) (*p2p.Server, error) {
	pkey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Generate private key failed with err: %v", err)
		return nil, err
	}

	cfg := p2p.Config{
		PrivateKey:      pkey,
		Name:            name,
		MaxPeers:        1,
		Protocols:       []p2p.Protocol{proto},
		EnableMsgEvents: true,
	}

	if port > 0 {
		cfg.ListenAddr = fmt.Sprintf(":%d", port)
	}
	srv := &p2p.Server{
		Config: cfg,
	}

	err = srv.Start()
	if err != nil {
		log.Printf("Start server failed with err: %v", err)
		return nil, err
	}

	return srv, nil
}

func connectToPeer(srv *p2p.Server, enode string) error {
	// Parsing the enode url
	node, err := discover.ParseNode(enode)
	if err != nil {
		log.Printf("Failed to parse enode url with err: %v", err)
		return err
	}

	// Connecting to the peer
	srv.AddPeer(node)

	return nil
}

func subscribeToEvents(srv *p2p.Server, communicated chan<- bool) {
	// Subscribing to the peer events
	peerEvent := make(chan *p2p.PeerEvent)
	eventSub := srv.SubscribeEvents(peerEvent)

	for {
		select {
		case event := <-peerEvent:
			if event.Type == p2p.PeerEventTypeMsgRecv {
				log.Println("Received message received notification")
				communicated <- true
			}
		case <-eventSub.Err():
			log.Println("subscription closed")

			// Closing the channel so that server gets stopped since
			// there won't be any more events coming in
			close(communicated)
		}
	}
}
