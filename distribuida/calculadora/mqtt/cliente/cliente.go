package main

import (
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

const qos = 1

func main() {

	// configurar cliente
	opts := MQTT.NewClientOptions()
	opts.AddBroker(shared.MQTTHost)
	opts.SetClientID("cliente")

	// criar novo cliente do broker
	client := MQTT.NewClient(opts)

	// conectar ao broker
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// desconectar cliente do broker
	defer client.Disconnect(250)

	// subscrever a um topico & usar um handler para receber as mensagens
	token = client.Subscribe(shared.MQTTReply, qos, receiveHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// loop
	for i := 0; i < 10; i++ {
		// cria a mensagem
		msg, err := json.Marshal(shared.Request{Op: "add", P1: 1, P2: 2})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//publicar a mensagem
		token := client.Publish(shared.MQTTRequest, qos, false, msg)
		token.Wait()
		if token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
		fmt.Printf("Mensagem Publicada: %s\n", msg)
		time.Sleep(time.Second)
	}
}

var receiveHandler MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	rep := shared.Reply{}
	err := json.Unmarshal(m.Payload(), &rep)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Recebida: ´%f´\n", rep.Result[0])
}
