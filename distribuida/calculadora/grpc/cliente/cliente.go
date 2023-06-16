package main

import (
	"distribuida/calculadora/grpc/calculadora"
	"distribuida/calculadora/grpc/fibonacci"
	"distribuida/calculadora/shared"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main() {
	var idx int32

	// Estabelece conexão com o servidor
	conn, err := grpc.Dial("localhost"+":"+strconv.Itoa(shared.GRPC_PORT), grpc.WithInsecure())
	shared.ChecaErro(err,"Não foi possível se conectar ao servidor")

	// fecha conexões
	defer conn.Close()

	// cria um proxy para o fibonacci
	fibo := fibonacci.NewFibonacciClient(conn)
	calc := calculadora.NewCalculadoraClient(conn)

	// contacta o servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx = 0; idx < shared.SAMPLE_SIZE; idx++ {
		// invoca operação remota
		fmt.Println(calc.Add(ctx, &calculadora.Request{Op: "add",P1:idx,P2:idx}))
		fmt.Println(fibo.Fibo(ctx, &fibonacci.Request{P1: idx}))
	}
}




