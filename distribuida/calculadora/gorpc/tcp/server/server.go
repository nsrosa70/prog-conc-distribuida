package main

import (
	calculadora "aulas/distribuida/calculadora/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

func servidor() {

	// cria instância da calculadora
	calculadora := new(calculadora.CalculadoraRPC)

	// cria um novo servidor RPC e registra a calculadora
	server := rpc.NewServer()
	err := server.RegisterName("Calculator", calculadora)
	shared.ChecaErro(err, "Não foi possível registrar a Calculadora no servidor...")

	// cria um listener TCP
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err, "Não foi possível criar um listener para a Calculadora...")
	defer func(ln net.Listener) {
		var err = ln.Close()
		shared.ChecaErro(err, "Não foi possível fechar o listener da Calculadora...")
	}(ln)

	// aguarda por invocações
	fmt.Println("Servidor está pronto (RPC-TCP)...")
	server.Accept(ln)
}

func main() {

	go servidor()

	_, _ = fmt.Scanln()
}
