package main

import (
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"math/rand"
	"time"
)

func main() {

	// gera nova seed
	rand.Seed(time.Now().UTC().UnixNano())

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// declara a fila para as respostas
	replyQueue, err := ch.QueueDeclare(
		shared.ResponseQueue,
		false,
		false,
		true,
		false,
		nil,
	)

	// cria consumer da fila de response
	msgs, err := ch.Consume(
		replyQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Falha ao registrar o consumer no broker")

	for i := 0; i < shared.SampleSize; i++ {
		// prepara mensagem
		msgRequest := shared.Request{Op: "add", P1: i + 1, P2: i + 1}
		msgRequestBytes, err := json.Marshal(msgRequest)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		correlationID := shared.RandomString(32)

		err = ch.Publish(
			"",
			shared.RequestQueue,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: correlationID,
				ReplyTo:       replyQueue.Name,
				Body:          msgRequestBytes,
			},
		)

		// recebe mensagem do consumer de mensageria
		m := <-msgs

		// deserializada e imprime mensagem na tela
		msgResponse := shared.Reply{}
		err = json.Unmarshal(m.Body, &msgResponse)
		shared.ChecaErro(err, "Erro na deserialização da resposta")
		fmt.Printf("%v(%v,%v)=%v\n", msgRequest.Op, msgRequest.P1, msgRequest.P2, msgResponse.Result[0])
	}
}
