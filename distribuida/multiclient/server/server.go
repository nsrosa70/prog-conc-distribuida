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
	EndMessage = "END"
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
	fmt.Println("Aguardando conexões dos cliente em " + ServerHost + ":" + ServerPort)
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Cliente conectado")

		// cria thread para o cliente
		go processRequestBytes(conn)
		//go processRequestJson(conn)
	}
}

func processRequestBytes(conn net.Conn) {
	var fromClient = make([]byte, 1024)

	for {
		// recebe dados
		mLen, err := conn.Read(fromClient)
		if err != nil {
			fmt.Println("Erro na leitura dos dados do cliente:", err.Error())
		}
		//fmt.Println("Dado recebido: ", string(fromClient[:mLen]))

		// envia resposta
		_, err = conn.Write([]byte(string(fromClient[:mLen])))

		if string(fromClient[:mLen]) == EndMessage {
			break
		}
	}

	// fecha conexão
	conn.Close()
}

func processRequestJson(conn net.Conn) {
	var fromClient string
	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)

	for {
		// recebe dados
		err := dec.Decode(&fromClient)
		if err != nil {
			fmt.Println("Erro na leitura dos dados do cliente:", err.Error())
		}
		fmt.Println("Dado recebido: ", fromClient)

		// envia resposta
		err = enc.Encode(fromClient)
		if err != nil {
			fmt.Println("Erro no envio dos dados para o cliente:", err.Error())
		}

		if fromClient == EndMessage {
			break
		}
	}

	// fecha conexão
	conn.Close()
}
