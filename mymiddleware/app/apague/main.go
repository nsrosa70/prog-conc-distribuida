package main

import (
	"fmt"
	"test/mymiddleware/infrastructure/crh"
	"test/mymiddleware/infrastructure/srh"
)

func main() {

	// Server
	s1 := srh.NewSRH("localhost", 1313)
	go func(id int, s *srh.SRH) {
		for {
			b := s.Receive()
			fmt.Printf("Servidor[%v]\n", id)
			s.Send(b)
		}
	}(1, s1)

	s2 := srh.NewSRH("localhost", 1314)
	go func(id int, s *srh.SRH) {
		for {
			b := s.Receive()
			fmt.Printf("Servidor[%v]\n", id)
			s.Send(b)
		}
	}(2, s2)

	// Client 1
	c1 := crh.NewCRH("localhost", 1313)
	go func(id int, c *crh.CRH) {
		for j := 0; j < 5; j++ {
			m := []byte{byte(j)}
			c.SendReceive(m)
			fmt.Printf("Client[%v] = %v\n", id, j)
		}
		fmt.Printf("Client %v finished\n", id)
	}(1, c1)

	// Client
	c2 := crh.NewCRH("localhost", 1314)
	go func(id int, c *crh.CRH) {
		for j := 0; j < 5; j++ {
			m := []byte{byte(j)}
			c.SendReceive(m)
			fmt.Printf("Client[%v] = %v\n", id, j)
		}
		fmt.Printf("Client %v finished\n", id)
	}(2, c2)

	fmt.Scanln()
}
