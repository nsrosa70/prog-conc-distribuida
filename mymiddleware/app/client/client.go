package main

import "aulas/mymiddleware/infrastructure/crh"

func main() {
	c := crh.NewCRH("localhost", 1313)

	b := []byte{1, 2, 3}
	c.SendReceive(b)
}
