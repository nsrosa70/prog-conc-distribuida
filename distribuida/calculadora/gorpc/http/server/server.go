package main

import (
	"aulas/distribuida/calculadora/gorpc/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

func main() {

	// 1: Criar a instância da messagingservice
	calculator := new(impl.Calculadora)

	// 2: Cria um novo servidor RPC e registrar a messagingservice
	server := rpc.NewServer()
	server.RegisterName("Calculadora", calculator)

	// 3: Associar um handler HTTP ao servidor
	server.HandleHTTP("/", "/debug")

	// 4: Criar um listener TCP
	l, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Servidor não inicializado")

	// 5: Aceitar e processar requisições remotas
	fmt.Printf("Servidor RPC pronto (RPC-HTTP) na porta %v...\n", shared.CalculatorPort)
	http.Serve(l, nil)
}
