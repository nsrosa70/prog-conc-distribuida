package shared

import (
	"log"
	"math/rand"
	"net"
	"strconv"
)

const MaxConnectionAttempts = 100

// MQTT
const MQTTHost = "mqtt://localhost:1883"
const MQTTTopic = "PubSub"
const MQTTRequest = "request"
const MQTTReply = "reply"

// Other configurations
const StatisticSample = 30
const SampleSize = 1000
const CalculatorPort = 4040
const GrpcPort = 5050
const NAMING_PORT = 1414
const MIOP_REQUEST = 1
const MIOP_REPLY = 2
const MAX_NUMBER_CLIENTS = 1
const RequestQueue = "request_queue"
const ResponseQueue = "response_queue"
const PubSubQueue = "pubsub"
const FanoutExchange = "fanout_exchange"
const DirectExchange = "direct_exchange"
const TopicExchange = "topic_exchange"
const HeadersExchange = "headers_exchange"
const RoutingKey = "routing_key"

const NamingPort = 1313
const CalculadoraPort = 1314
const FibonacciPort = 1315
const MessagingPort = 1316
const LocalHost = "localhost"
const CallBackPort = 1317

const MaxQueueSize = 10000

type Message struct {
	Payload string
}

type Args struct {
	A, B int
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

type Invocation struct {
	Ior     IOR
	Request Request
}

type Termination struct {
	Rep Reply
}

type IOR struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

func NewIOR(h string, p int) IOR {
	return IOR{Host: h, Port: p, Id: 0, TypeName: ""}
}

type Request struct {
	Op     string
	Params []interface{}
}

type Reply struct {
	Result []interface{}
}

func AvailablePort() int {
	r := 1024

	for port := 1024; port < 2000; port++ {
		ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))

		if err == nil {
			ln.Close()
			r = port
		}
	}
	return r
}
