// rpc_server.go

package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"test/distribuida/calculadora/gorpc/impl"
	"test/shared"
)

func main() {
	// 1: Criar instância da calculadora.
	mathService := new(impl.Calculadora)

	// 2: Registrar a instância da calculadora no RPC
	server := rpc.NewServer()
	err := server.Register(mathService)
	shared.ChecaErro(err, "Erro ao registrar a calculadora")

	// 3: Criar listener para as conexões remotas
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Erro ao iniciar o listener")
	defer listener.Close()

	fmt.Printf("Servidor RPC pronto (RPC-TCP) na porta %v...\n", shared.CalculatorPort)

	// 4: Aceitar e processar requisições remotas
	server.Accept(listener)
}
