// socket-client project main.go
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	ServerHost     = "localhost"
	ServerPort     = "1313"
	ServerType     = "tcp"
	SampleSize     = 30
	NumberRequests = 5
	EndMessage     = 0
)

func main() {

	for i := 0; i < SampleSize; i++ {

		// estabelece conexão
		conn, err := net.Dial(ServerType, ServerHost+":"+ServerPort)
		if err != nil {
			panic(err)
		}

		// envia request/recebe resposta
		comServer(conn)

		// fecha conexão
		defer conn.Close()
	}
}

func comServer(conn net.Conn) {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)
	var fromServer, toServer int

	for i := 0; i < NumberRequests; i++ {

		// envia mensagem para o consumer
		toServer = 8
		err := enc.Encode(toServer)
		if err != nil {
			fmt.Println("Erro no envio dos dados do consumer:", err.Error())
		}

		// recebe resposta do consumer
		err = dec.Decode(&fromServer)
		if err != nil {
			fmt.Println("Erro no recebimento dos dados do consumer:", err.Error())
		}
		fmt.Printf("Fibonacci (%v) = %v\n", toServer, fromServer)
	}

	// envia mensagem de fim de dados
	err := enc.Encode(EndMessage)
	if err != nil {
		fmt.Println("Erro no envio dos dados do consumer:", err.Error())
	}
}
