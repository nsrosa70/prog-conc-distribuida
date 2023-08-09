package main

import (
	"aulas/distribuida/shared"
	"fmt"
	"net/rpc"
	"strconv"
)

func main() {

	var reply int

	// conecta ao consumer
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "O Servidor não está pronto")

	// prepara request
	args := shared.Args{A: 1, B: 2}

	// envia request e recebe resposta
	client.Call("Calculator.Add", args, &reply)

	fmt.Printf("Add(%v,%v) = %v \n", args.A, args.B, reply)
}
