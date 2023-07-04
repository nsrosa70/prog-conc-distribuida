// socket-client project main.go
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	ServerHost     = "localhost"
	ServerPort     = "1313"
	ServerType     = "tcp"
	SampleSize     = 30
	NumberRequests = 1000
	EndMessage     = "END"
)

func main() {

	for i := 0; i < SampleSize; i++ {

		// estabelece conexão
		conn, err := net.Dial(ServerType, ServerHost+":"+ServerPort)
		if err != nil {
			panic(err)
		}

		// envia dado/recebe resposta
		t1 := time.Now()
		comServerBytes(conn)
		//comServerJson(conn)
		fmt.Println(time.Now().Sub(t1).Milliseconds())

		// fecha conexão
		defer conn.Close()
	}
}

func comServerBytes(conn net.Conn) {
	fromServer := make([]byte, 1024)
	toServer := ""

	for i := 0; i < NumberRequests; i++ {

		// envia mensagem
		toServer = "Mensagem #" + strconv.Itoa(i)
		_, err := conn.Write([]byte(toServer))
		if err != nil {
			fmt.Println("Erro no envio dos dados do servidor:", err.Error())
		}

		// recebe resposta do servidor
		//mLen, err := conn.Read(fromServer)
		_, err = conn.Read(fromServer)
		if err != nil {
			fmt.Println("Erro no recebimento dos dados do servidor:", err.Error())
		}
		//fmt.Println("Dado: ", string(fromServer[:mLen]))
	}
}

func comServerJson(conn net.Conn) {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)
	fromServer := ""
	toServer := ""

	for i := 0; i < NumberRequests; i++ {

		// envia mensagem para o servidor
		toServer = "Message #" + strconv.Itoa(i)
		err := enc.Encode(toServer)
		if err != nil {
			fmt.Println("Erro no envio dos dados do servidor:", err.Error())
		}

		// recebe resposta do servidor
		err = dec.Decode(&fromServer)
		if err != nil {
			fmt.Println("Erro no recebimento dos dados do servidor:", err.Error())
		}
		//fmt.Println("Dado: ", fromServer)
	}

	// envia mensagem de fim de dados
	err := enc.Encode(EndMessage)
	if err != nil {
		fmt.Println("Erro no envio dos dados do servidor:", err.Error())
	}
}
