package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	id := 0
	for {
		id++
		go func(id int) {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			defer conn.Close()

			//fmt.Printf("Client %d Connected to server\n", *id)
			nMessages := rand.Intn(1000)

			for i := 0; i < nMessages; i++ {
				message := strconv.Itoa(rand.Intn(100000))
				_, err := conn.Write([]byte(message + "\n")) // Send message
				if err != nil {
					fmt.Println("Connection closed by server.")
					return
				}

				response, _ := bufio.NewReader(conn).ReadString('\n') // Read response
				response = response
				fmt.Printf("Client [%d]: %v", id, response)
			}
		}(id)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Scanln()
}
