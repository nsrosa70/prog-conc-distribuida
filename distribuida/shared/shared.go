package shared

import (
	"log"
	"math/rand"
	"net"
	"strconv"
)

const SAMPLE_SIZE = 10
const CALCULATOR_PORT = 4040
const FIBONACCI_PORT = 3030
const GRPC_PORT = 5050
const NAMING_PORT = 1414
const MIOP_REQUEST = 1
const MIOP_REPLY = 2
const MAX_NUMBER_CLIENTS = 1

type Message struct {
	Payload string
}
type Request struct {
	Op string
	P1 int
	P2 int
}

type Reply struct {
	Result []interface{}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}

func FindNextAvailablePort() int { // TCP only
	i := 3000

	for i = 3000; i < 4000; i++ {
		port := strconv.Itoa(i)
		ln, err := net.Listen("tcp", ":"+port)

		if err == nil {
			ln.Close()
			break
		}
	}
	return i
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}