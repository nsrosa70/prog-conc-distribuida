package main

import (
	"aulas/mymiddleware/infrastructure/srh"
	"fmt"
)

func main() {
	s := srh.NewSRH("localhost", 1313)
	for {
		b := s.Receive()
		s.Send(b)
		fmt.Println(b)
	}
}
