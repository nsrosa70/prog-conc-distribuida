package main

import (
	"aulas/distribuida/calculadora/shared"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func clientRPCTCPPerformance() {
	var reply int
	times := []time.Duration{}
	var SAMPLE_SIZE = 1000

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// loop
	for i := 0; i < SAMPLE_SIZE; i++ {

		// prepara request & start time
		t1 := time.Now()
		args := shared.Args{A: i, B: i}

		// invoca operação remota
		client.Call("Calculator.Add", args, &reply)

		// stop time
		times[i] = time.Now().Sub(t1)
	}
	totalTime := time.Duration(0)
	for i := range times {
		totalTime += times[i]
	}
	fmt.Printf("Total Duration: %v [%v]", totalTime, shared.SAMPLE_SIZE)
}

func Cliente() {
	var reply int

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Fatal(err)
	}

	defer func(client *rpc.Client) {
		var err = client.Close()
		if err != nil {

		}
	}(client)

	// invoca operação remota
	args := shared.Args{A: 1, B: 2}
	err = client.Call("Calculator.Add", args, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v + %v = %v \n", args.A, args.B, reply)
}

func main() {

	go Cliente()

	fmt.Scanln()
}
