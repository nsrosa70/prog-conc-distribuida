package main

import (
	calculadora "aulas/distribuida/calculadora/impl"
	"aulas/distribuida/shared"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

const qos = 1

func main() {

	// configurar cliente
	opts := MQTT.NewClientOptions()
	opts.AddBroker(shared.MQTTHost)
	opts.SetClientID("subscriber 1")
	//opts.DefaultPublishHandler = receiveHandler

	// criar novo cliente
	client := MQTT.NewClient(opts)

	// conectar ao broker
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// desconectar ao broker
	defer client.Disconnect(250)

	// subscrever a um topico & usar um handler para receber as mensagens
	token = client.Subscribe(shared.MQTTRequest, qos, receiveHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Println("Consumidor iniciado...")
	fmt.Scanln()
}

var receiveHandler MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	req := shared.Request{}
	err := json.Unmarshal(m.Payload(), &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := calculadora.Calculadora{}.InvocaCalculadora(req)
	rep, err := json.Marshal(shared.Reply{[]interface{}{r}})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Marshal(rep)
	token := c.Publish(shared.MQTTReply, qos, false, rep)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	fmt.Printf("Recebida: ´%s´ Enviada: ´%s´\n", req, rep)
}
