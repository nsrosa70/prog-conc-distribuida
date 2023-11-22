package main

import (
	impl "aulas/distribuida/calculadora/implGoRPC"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

func servidor() {

	// cria instância da calculadora
	calculadora := new(impl.CalculadoraGoRPC)

	// cria um novo servidor RPC e registra a calculadora
	server := rpc.NewServer()
	err := server.RegisterName("Calculator", calculadora)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// cria um listenet TCP
	ln, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(ln net.Listener) {
		var err = ln.Close()
		if err != nil {

		}
	}(ln)

	// aguarda por invocações
	fmt.Println("Servidor está pronto ...")
	server.Accept(ln)
}

func main() {

	go servidor()

	_, _ = fmt.Scanln()
}
