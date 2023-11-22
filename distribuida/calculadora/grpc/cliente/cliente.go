package main

import (
	gen1 "aulas/distribuida/calculadora/grpc/proto"
	gen2 "aulas/distribuida/fibonacci/proto"
	"aulas/distribuida/shared"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"time"
)

func main() {

	// Estabelece conexão com o servidor
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	endPoint := "localhost" + ":" + strconv.Itoa(shared.GrpcPort)
	conn, err := grpc.Dial(endPoint, opt)
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor em"+endPoint)

	// fecha conexões
	defer conn.Close()

	// cria um proxy para a calculadora
	calc := gen1.NewCalculadoraClient(conn)

	// cria um contexto para execução remota
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var i int32
	for i = 0; i < shared.SampleSize; i++ {
		// invoca operação remota
		reqCalc := gen1.Request{P1: i, P2: i}
		reqFibo := gen2.RequestFibo{P1: i}
		repCalc, err := calc.Add(ctx, &reqCalc)
		shared.ChecaErro(err, "Erro ao invocar a operação remota.")
		repFibo, err := calc.Add(ctx, &reqCalc)
		shared.ChecaErro(err, "Erro ao invocar a operação remota.")
		fmt.Printf("Add(%v,%v)=%v Fibo(%v)=%v\n", reqCalc.P1, reqCalc.P2, repCalc.N, reqFibo.P1, repFibo.N)
	}
}
