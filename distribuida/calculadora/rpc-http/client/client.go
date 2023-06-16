package main
import (
	"fmt"
	"net/rpc"
	"shared"
	"strconv"
	"time"
)

func main() {

	var reply int
	// connect to servidor
	client, err := rpc.DialHTTP("tcp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	shared.ChecaErro(err,"O Servidor não está pronto")


	// make requests
	//fmt.Println("Client started execution... ")
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// prepara request
		args := shared.Args{A: i, B: i}

		// envia request e recebe resposta
		client.Call("Calculator.Add", args, &reply)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Println(x)
	}
	//elapsed := time.Since(start)
	//fmt.Printf("Tempo: %s \n", elapsed)
}
