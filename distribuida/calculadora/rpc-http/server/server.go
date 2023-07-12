package main

import (
	"aulas/distribuida/calculadora/impl"
	"aulas/distribuida/calculadora/shared"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {

	// create new instance of calculator
	calculator := new(impl.CalculadoraRPC)

	// create new rpc servidor
	server := rpc.NewServer()
	server.RegisterName("Calculator", calculator)

	// associate a http handler to servidor
	server.HandleHTTP("/", "/debug")

	// create tcp listen
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err, "Servidor não inicializado")

	// wait for calls
	fmt.Println("Servidor pronto (RPC-HTTP) ...\n")
	http.Serve(l, nil)
}
