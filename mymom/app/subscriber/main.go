package main

import (
	messagingproxy "test/mymom/services/messagingservice/proxy"
	"test/shared"
)

func main() {
	// Obtain messagingservice proxies
	messaging := messagingproxy.New(shared.LocalHost, shared.MessagingPort)

	messaging = messaging
	// Receive messages
	for {
		// receive
	}
}
