package main

import (
	"distribuida/calculadora/impl"
	"distribuida/calculadora/shared"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func CalculatorServerTCP() {

	// return an address of a TCP end point
	r,err := net.ResolveTCPAddr("tcp",":"+strconv.Itoa(shared.CALCULATOR_PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Listen on tcp port
	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Server is ready to accept connections (TCP)...")

	for {
		// Accept new connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Handle connection
		go HandleTCP(conn)
	}
}

func HandleTCP(conn net.Conn) {
	var msgFromClient shared.Request

	// Close connection
	defer conn.Close()

	// Create coder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		// Receive request
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			//conn.Close()
			// no further requests
			break
		}

		// Process request
		r := impl.Calculadora{}.InvocaCalculadora(msgFromClient)

		// Create response
		msgToClient := shared.Reply{[]interface{}{r}}

		// Serialise and send response to client
		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			os.Exit(0)
			break
		}
	}

}

func main() {

	go CalculatorServerTCP()

	fmt.Scanln()
}

