package main

import (
	"distribuida/calculadora/shared"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func CalculatorClientTCP() {
	var response shared.Reply

	port := "1313"

	// return an address of a TCP end point
	r, err := net.ResolveTCPAddr("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Connect to server
	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Close connection
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Create enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	/*
		// warm-up
		for i := 0; i < 1000; i++ { //N = 1000, 10000, 1000000

			// Prepare request
			msgToServer := shared.Request{"add", 3, 1}

			// Serialise and send request to consumer
			err = jsonEncoder.Encode(msgToServer)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			// Receive response from server
			err = jsonDecoder.Decode(&response)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			//fmt.Printf("%s(%d,%d) = %.0f \n",msgToServer.Op,msgToServer.P1,msgToServer.P2,response.Result[0].(float64))
		}
	*/
	// Measurement

	//var t2 time.Duration
	//t2 = 0

	//tt := [100000]time.Duration{}

	//for i := 0; i < shared.SAMPLE_SIZE; i++ { //N = 1000, 10000, 1000000
	for i := 0; i < 10; i++ { //N = 1000, 10000, 1000000

		// Prepare request
		msgToServer := shared.Request{"add", 3, 1}

		//t1 := time.Now()

		// Serialise and send request to consumer
		err = jsonEncoder.Encode(msgToServer)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive response from consumer
		err = jsonDecoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		//t2 = t2 + time.Now().Sub(t1)

		fmt.Printf("%s(%d,%d) = %.0f \n", msgToServer.Op, msgToServer.P1, msgToServer.P2, response.Result[0].(float64))
	}

	//fmt.Println(t2/shared.SAMPLE_SIZE)
}

func main() {

	go CalculatorClientTCP()

	_, _ = fmt.Scanln()
}
