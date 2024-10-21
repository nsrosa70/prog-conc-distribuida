// rpc_client.go

package main

import (
	"fmt"
	"net/rpc"
	"strconv"
	"test/distribuida/calculadora/gorpc/impl"
	"test/shared"
	"time"
)

func main() {
	ClientePerf()
}

func Cliente() {
	// 1: Conectar ao servidor RPC - host/porta
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CalculadoraPort))
	shared.ChecaErro(err, "Erro ao conectar ao servidor")
	defer client.Close()

	// 2: Invocar a operação remota
	req := impl.Request{P1: 10, P2: 20}
	rep := impl.Reply{}
	err = client.Call("Calculadora.Add", req, &rep)
	shared.ChecaErro(err, "Erro na invocação remota")

	// 3: Imprimir o resultado
	fmt.Printf("Add(%v,%v) = %v \n", req.P1, req.P2, rep.R)
}

func ClientePerf() {
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(shared.CalculatorPort))
	shared.ChecaErro(err, "Erro ao conectar ao servidor")
	defer client.Close()

	req := impl.Request{P1: 1, P2: 2}
	rep := impl.Reply{}
	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			err = client.Call("Calculadora.Add", req, &rep)
			shared.ChecaErro(err, "Erro na invocação da Calculadora remota...")
		}
		//fmt.Printf("tcp;%v\n", time.Now().Sub(t1).Milliseconds())
		fmt.Println(i, ";", time.Now().Sub(t1).Milliseconds())

	}
}
