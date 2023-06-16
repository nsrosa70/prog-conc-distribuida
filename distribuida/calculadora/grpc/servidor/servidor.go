package main

import (
	"distribuida/calculadora/grpc/calculadora"
	"distribuida/calculadora/grpc/fibonacci"
	"distribuida/calculadora/impl"
	"distribuida/calculadora/shared"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

func main() {

	conn, err := net.Listen("tcp", ":"+strconv.Itoa(shared.GRPC_PORT))
	shared.ChecaErro(err,"Não foi possível criar o listener")

	server := grpc.NewServer()

	calculadora.RegisterCalculadoraServer(server,&impl.CalculadoraGRPC{})
	fibonacci.RegisterFibonacciServer(server,&impl.Fibonacci{})

	reflection.Register(server)

	fmt.Println("Servidor pronto ...")

	err = server.Serve(conn)
	shared.ChecaErro(err,"Falha ao iniciar servidor")
}



