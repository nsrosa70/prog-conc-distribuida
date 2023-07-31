package main

import (
	"aulas/distribuida/shared"
	"fmt"
	"log"
	"net/rpc"
	"strconv"
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

	// conecta ao servidor (Calculadora)
	clientCalc, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err, "Não foi possível estabelecer uma conexão TCP com o servidor da Calculadora...")

	defer func(clientCalc *rpc.Client) {
		var err = clientCalc.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor da Calculadora...")
	}(clientCalc)

	// invoca operação remota da calculadora
	args := shared.Args{A: 3, B: 4}
	err = clientCalc.Call("Calculator.Add", args, &reply)
	shared.ChecaErro(err, "Erro na invocação da Calculadora remota...")

	fmt.Printf("Add(%v,%v) = %v \n", args.A, args.B, reply)
}

func main() {

	go Cliente()

	fmt.Scanln()
}
