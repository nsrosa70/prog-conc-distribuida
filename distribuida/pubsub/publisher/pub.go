package main

import (
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	//go defaultPublihser(&wg)
	//go directPublihser(&wg)
	//go fanoutPublihser(&wg)
	go topicPublisher(&wg)
	//go headersPublisher(&wg)

	wg.Wait()

}

func defaultPublihser(wg *sync.WaitGroup) {
	defer wg.Done()

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// envia mensagens
	for i := 0; i < shared.SampleSize; i++ {

		// prepara mensagem
		msg := shared.Message{Payload: "Mensagem-" + strconv.Itoa(i)}
		msgBytes, err := json.Marshal(msg)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica mensagem
		err = ch.Publish("", // exchange
			shared.PubSubQueue, // routing key
			false,
			false,
			amqp.Publishing{ContentType: "text/plain", Body: msgBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o consumer de mensageria")

		fmt.Printf("Publisher[Default]: %v \n", msg.Payload)
	}
}

func directPublihser(wg *sync.WaitGroup) {
	defer wg.Done()

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// declara o tipo de exchange
	err = ch.ExchangeDeclare(shared.DirectExchange,
		"direct",
		false,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Erro ao criar exchange")

	for i := 0; i < shared.SampleSize; i++ {

		// prepara mensagem
		msg := shared.Message{Payload: "Mensagem-" + strconv.Itoa(i)}
		msgBytes, err := json.Marshal(msg)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica mensagem
		err = ch.Publish(shared.DirectExchange, // exchange
			shared.RoutingKey, // routing key
			false,
			false,
			amqp.Publishing{ContentType: "text/plain", Body: msgBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o consumer de mensageria")

		fmt.Printf("Publisher[Direct]: %v \n", msg.Payload)
	}
}

func fanoutPublihser(wg *sync.WaitGroup) {
	defer wg.Done()

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// declara o tipo de exchange
	err = ch.ExchangeDeclare(shared.FanoutExchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Erro ao criar exchange")

	for i := 0; i < shared.SampleSize; i++ {

		// prepara mensagem
		msg := shared.Message{Payload: "Mensagem-" + strconv.Itoa(i)}
		msgBytes, err := json.Marshal(msg)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica mensagem
		err = ch.Publish(shared.FanoutExchange, // exchange
			"", // routing key
			false,
			false,
			amqp.Publishing{ContentType: "text/plain", Body: msgBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o consumer de mensageria")

		fmt.Printf("Publisher[Fanout]: %v \n", msg.Payload)
	}
}

func topicPublisher(wg *sync.WaitGroup) {
	defer wg.Done()

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// declara o tipo de exchange
	err = ch.ExchangeDeclare(shared.TopicExchange,
		"topic",
		false,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Erro ao criar exchange")

	for i := 0; i < shared.SampleSize; i++ {

		// prepara mensagem
		msg := shared.Message{Payload: "Mensagem-" + strconv.Itoa(i)}
		msgBytes, err := json.Marshal(msg)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica mensagem
		err = ch.Publish(shared.TopicExchange, // exchange
			"minha.chave", // routing key
			false,
			false,
			amqp.Publishing{ContentType: "text/plain", Body: msgBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o consumer de mensageria")

		fmt.Printf("Publisher[Topic]: %v \n", msg.Payload)
	}
}

func headersPublisher(wg *sync.WaitGroup) {
	defer wg.Done()

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao consumer de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o consumer de mensageria")
	defer ch.Close()

	// declara o tipo de exchange
	err = ch.ExchangeDeclare(shared.HeadersExchange,
		"headers",
		false,
		false,
		false,
		false,
		nil)
	shared.ChecaErro(err, "Erro ao criar exchange")

	for i := 0; i < shared.SampleSize; i++ {

		// prepara mensagem
		msg := shared.Message{Payload: "Mensagem-" + strconv.Itoa(i)}
		msgBytes, err := json.Marshal(msg)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		headers := amqp.Table{
			"at1": "xxx",
			"at2": "oi2",
		}
		// publica mensagem
		err = ch.Publish(shared.HeadersExchange,
			// exchange
			"", // routing key
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msgBytes,
				Headers:     headers})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o consumer de mensageria")

		fmt.Printf("Publisher[Topic]: %v \n", msg.Payload)
	}
}
