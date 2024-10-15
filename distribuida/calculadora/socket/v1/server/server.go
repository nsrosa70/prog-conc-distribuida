package main

import (
	"aulas/distribuida/calculadora/impl"
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	CalculatorServerTCP()

	//CalculatorServerUDP()

	fmt.Scanln()
}

func CalculatorServerTCP() {

	//  define o endpoint do servidor TCP
	r, err := net.ResolveTCPAddr("tcp", "localhost:1314")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// cria um listener TCP
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Servidor TCP aguardando conexões na porta 1314...")

	for {
		// aguarda/aceita conexão
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// processa requests da conexão
		go HandleTCPConnection(conn)
	}
}

func HandleTCPConnection(conn net.Conn) {
	var msgFromClient shared.Request

	// Close connection
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	// Cria coder/decoder JSON
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		// recebe & unmarshall requests do cliente
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		// processa request
		r := impl.Calculadora{}.InvocaCalculadora(msgFromClient)

		// cria resposta
		msgToClient := shared.Reply{[]interface{}{r}}

		// serializa & envia resposta para o cliente
		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func CalculatorServerUDP() {
	msgFromClient := make([]byte, 1024)

	// resolve subscriber address
	addr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(shared.CalculatorPort))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// listen on udp port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// close conn
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fmt.Println("Server UDP is ready to accept requests at port", shared.CalculatorPort, "...")

	for {
		// receive request
		n, addr, err := conn.ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// handle request
		HandleUDPRequest(conn, msgFromClient, n, addr)
	}
}

func HandleUDPRequest(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var request shared.Request

	//unmarshall request
	err := json.Unmarshal(msgFromClient[:n], &request)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// process request
	r := impl.Calculadora{}.InvocaCalculadora(request)

	// create response
	rep := shared.Reply{[]interface{}{r}}

	// serialise response
	msgToClient, err = json.Marshal(rep)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// send response
	_, err = conn.WriteTo(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
