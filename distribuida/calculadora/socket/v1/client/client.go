package main

import (
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	//t1 := time.Now()
	CalculatorClientTCP(10)
	//CalculatorClientUDP(n)
	//tTotal := time.Now().Sub(t1)

	//fmt.Println(tTotal.Nanoseconds()/1000000)
	//CalculatorClientUDP(n)
}

func CalculatorClientTCP(n int) {
	var response shared.Reply

	// retorna o endereço do endpoint
	r, err := net.ResolveTCPAddr("tcp", "localhost:1314")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	/// connecta ao consumer (sem definir uma porta local)
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha conexõa
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// cria enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < n; i++ {

		// prepara request
		msgToServer := shared.Request{"add", i, i}

		// serializa e envia request para o consumer
		err = jsonEncoder.Encode(msgToServer)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// recebe resposta do consumer
		err = jsonDecoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(response)
	}
}

func CalculatorClientUDP(n int) {
	var response shared.Reply

	// resolve server address
	addr, err := net.ResolveUDPAddr("udp", "localhost:"+strconv.Itoa(shared.CalculatorPort))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// connect to server -- does not create a connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// create coder/decoder
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	// Close connection
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	for i := 0; i < shared.SampleSize; i++ {
		// Create request
		request := shared.Request{Op: "add", P1: n, P2: n}

		// Serialise and send request
		err = encoder.Encode(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive response from consumer
		err = decoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		//fmt.Printf("%s(%d,%d) = %.0f \n", request.Op, request.P1, request.P2, response.Result[0])
	}
}
