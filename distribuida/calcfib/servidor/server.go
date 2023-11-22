package main

import (
	calculadora "aulas/distribuida/calculadora/impl"
	fibonacci "aulas/distribuida/fibonacci/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func servidor() {

	// cria instância da calculadora/fibonacci
	calculadora := new(calculadora.CalculadoraRPC)
	fibonacci := new(fibonacci.FibonacciRPC)

	// cria um novo servidor e registra a calculadora
	server := rpc.NewServer()
	err := server.RegisterName("Calculator", calculadora)
	shared.ChecaErro(err, "Não foi possível registrar a Calculadora no servidor...")

	err = server.RegisterName("Fibonacci", fibonacci)
	shared.ChecaErro(err, "Não foi possível registrar o Fibonacci no servidor...")

	// cria um listener TCP (Calculadora)
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Listen TCP da Calculadora não pode ser criado...")
	defer func(ln net.Listener) {
		var err = ln.Close()
		shared.ChecaErro(err, "Listen TCP da Calculadora não pode ser fechado...")
	}(ln)

	// // cria um listener TCP (fibonacci)
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.FibonacciPort))
	shared.ChecaErro(err, "Listen TCP do Fibonacci não pode ser criado...")

	// associa um handler HTTP ao servidor (Fibonacci)
	server.HandleHTTP("/", "/debug")

	// aguarda por invocações
	fmt.Println("Servidor está pronto (RPC-TCP/HTTP)...")
	go server.Accept(ln) // calculadora
	http.Serve(l, nil)   // fibonacci
}

func main() {

	go servidor()

	_, _ = fmt.Scanln()
}
