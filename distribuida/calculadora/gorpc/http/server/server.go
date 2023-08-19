package main

import (
	"aulas/distribuida/calculadora/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {

	// cria instância da calculadora
	calculator := new(impl.CalculadoraRPC)

	// cria um novo servidor RPC e registra a calculadora
	server := rpc.NewServer()
	server.RegisterName("Calculator", calculator)

	// associa um handler HTTP ao servidor
	server.HandleHTTP("/", "/debug")

	// // cria um listener TCP
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Servidor não inicializado")

	// aguarda por invocações
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}
