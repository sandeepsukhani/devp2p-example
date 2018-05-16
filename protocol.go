package main

// p2p Protocol definition for sending and receiving a message
import (
	"fmt"

	"github.com/ethereum/go-ethereum/p2p"
)

var (
	proto = p2p.Protocol{
		Name:    "ping",
		Version: 1,
		Length:  1,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			message := "ping"

			// Sending the message to connected peer
			err := p2p.Send(rw, 0, message)
			if err != nil {
				return fmt.Errorf("Send message fail: %v", err)
			}
			fmt.Println("sending message", message)

			// Receiving the message from connected peer
			received, err := rw.ReadMsg()
			if err != nil {
				return fmt.Errorf("Receive message fail: %v", err)
			}

			var myMessage string
			err = received.Decode(&myMessage)

			fmt.Println("received message", string(myMessage))

			return nil
		},
	}
)
