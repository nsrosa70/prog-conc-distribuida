// Para Quickstart do gRPC, acesse https://grpc.io/docs/languages/go/quickstart/
package main

import (
	fibonacci "aulas/distribuida/fibonacci/impl"
	gen2 "aulas/distribuida/fibonacci/proto"
	gen1 "aulas/distribuida/messagingservice/grpc/proto"
	calculadora "aulas/distribuida/messagingservice/impl"
	"aulas/distribuida/shared"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

func main() {

	// Cria listener
	endPoint := "localhost:" + strconv.Itoa(shared.GrpcPort)
	conn, err := net.Listen("tcp", endPoint)
	shared.ChecaErro(err, "Não foi possível criar o listener")

	// Cria um gRPC Server (“serviço de nomes” + servidor)”
	server := grpc.NewServer()

	// Registra a “Calculadora"/"Fibonacci" no servidor de nomes
	gen1.RegisterCalculadoraServer(server, &calculadora.CalculadoraRPC{})
	gen2.RegisterFibonacciServer(server, &fibonacci.FibonacciRPC{})
	reflection.Register(server)

	fmt.Println("Servidor pronto ...")

	// Inicia servidor para atender requisções
	err = server.Serve(conn)
	shared.ChecaErro(err, "Falha ao iniciar servidor")
}
