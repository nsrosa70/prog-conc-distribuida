package main

import (
	"distribuida/calculadora/impl"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"shared"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err,"Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	ch, err := conn.Channel()
	shared.ChecaErro(err,"Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// declaração de filas
	requestQueue, err := ch.QueueDeclare("request", false, false, false,
		false, nil)
	shared.ChecaErro(err,"Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare("response", false, false, false,
		false, nil)
	shared.ChecaErro(err,"Não foi possível criar a fila no servidor de mensageria")

	// prepara o recebimento de mensagens do clientserver
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil)
	shared.ChecaErro(err,"Falha ao registrar o consumidor do servidor de mensageria")

	fmt.Println("Servidor pronto...")
	for d := range msgsFromClient {

		// recebe request
		msgRequest := shared.Request{}
		err := json.Unmarshal(d.Body, &msgRequest)
		shared.ChecaErro(err,"Falha ao desserializar a mensagem")

		// processa request
		r := impl.Calculadora{}.InvocaCalculadora(msgRequest)

		// prepara resposta
		replyMsg := shared.Reply{Result: []interface{}{r}}
		replyMsgBytes, err := json.Marshal(replyMsg)
		shared.ChecaErro(err,"Não foi possível criar a fila no servidor de mensageria")
		if err != nil {
			log.Fatalf("%s: %s", "Falha ao serializar mensagem", err)
		}

		// publica resposta
		err = ch.Publish("", replyQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes)})
		shared.ChecaErro(err,"Falha ao enviar a mensagem para o servidor de mensageria")
	}
}
