package main

import (
	"aulas/distribuida/shared"
	"fmt"
	"net/rpc"
	"strconv"
)

func Cliente() {
	var reply int

	// conecta ao servidor (Calculadora)
	clientCalc, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Conexão TCP para o servidor da Calucladora não pode ser criada...")

	defer func(clientCalc *rpc.Client) {
		var err = clientCalc.Close()
		shared.ChecaErro(err, "Não foi possível fechar a Conexão TCP para o servidor da Calculadora...")
	}(clientCalc)

	// conecta ao servidor (Fibonacci)
	clientFibo, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.FibonacciPort))
	shared.ChecaErro(err, "Conexão TCP para o servidor Fibonacci não pode ser criada...")

	defer func(clientFibo *rpc.Client) {
		var err = clientFibo.Close()
		shared.ChecaErro(err, "Não foi possível fechar a Conexão TCP para o servidor do Fibonacci...")
	}(clientFibo)

	// invoca operação remota da messagingservice
	args := shared.Args{A: 3, B: 4}
	err = clientCalc.Call("Calculator.Add", args, &reply)
	shared.ChecaErro(err, "Erro na invocação remora da Calucladora...")

	fmt.Printf("Add(%v,%v) = %v \n", args.A, args.B, reply)

	// invoca operação remota do fibonacci
	n := 10
	err = clientFibo.Call("Fibonacci.Fibo", n, &reply)
	shared.ChecaErro(err, "Erro na invocação remora do Fibonacci...")

	fmt.Printf("Fibo(%v) = %v \n", n, reply)

}

func main() {

	go Cliente()

	fmt.Scanln()
}
