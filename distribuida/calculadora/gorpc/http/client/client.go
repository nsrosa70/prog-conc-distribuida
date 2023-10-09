package main

import (
	"aulas/distribuida/calculadora/gorpc/impl"
	"aulas/distribuida/shared"
	"fmt"
	"net/rpc"
	"strconv"
	"time"
)

func main() {
	ClientePerf()
}

func Cliente() {

	// 1: Conectar ao servidor (Calculadora)
	clientCalc, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "O Servidor não está pronto")
	defer func(clientCalc *rpc.Client) {
		var err = clientCalc.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor da Calculadora...")
	}(clientCalc)

	// 2: Invocar operação remota da calculadora
	req := impl.Request{P1: 10, P2: 20}
	rep := impl.Reply{}
	err = clientCalc.Call("Calculadora.Add", req, &rep)
	shared.ChecaErro(err, "Erro na invocação da Calculadora remota...")

	// 3: Imprimir o resultado
	fmt.Printf("Add(%v,%v) = %v \n", req.P1, req.P2, rep.R)
}

func ClientePerf() {

	// 1: Conectar ao servidor (Calculadora)
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "O Servidor não está pronto")
	defer func(clientCalc *rpc.Client) {
		var err = clientCalc.Close()
		shared.ChecaErro(err, "Não foi possível fechar a conexão TCP com o servidor da Calculadora...")
	}(client)

	req := impl.Request{P1: 10, P2: 20}
	rep := impl.Reply{}
	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			err = client.Call("Calculadora.Add", req, &rep)
			shared.ChecaErro(err, "Erro na invocação da Calculadora remota...")
		}
		fmt.Printf("http;%v\n", time.Now().Sub(t1).Milliseconds())
	}
}
