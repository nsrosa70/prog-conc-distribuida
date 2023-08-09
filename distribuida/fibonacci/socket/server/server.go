// socket-server project main.go
// developer.com/languages/intro-socket-programming-go/
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	ServerHost = "localhost"
	ServerPort = "1313"
	ServerType = "tcp"
	EndMessage = 0
)

func main() {

	// cria listener
	fmt.Println("Servidor em execução...")
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Erro na escuta por conexões:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	// aguarda conexões
	fmt.Println("Aguardando conexões dos publisher em " + ServerHost + ":" + ServerPort)
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Cliente conectado")

		// cria thread para o publisher
		go processRequest(conn)
	}
}

func processRequest(conn net.Conn) {
	var fromClient, toClient int
	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)

	for {
		// recebe dados
		err := dec.Decode(&fromClient)
		if err != nil {
			fmt.Println("Erro na leitura dos dados do publisher:", err.Error())
		}

		// envia resposta
		toClient = fibonacci(fromClient)
		err = enc.Encode(toClient)
		if err != nil {
			fmt.Println("Erro no envio dos dados para o publisher:", err.Error())
		}

		if fromClient == EndMessage {
			break
		}
	}

	// fecha conexão
	conn.Close()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
