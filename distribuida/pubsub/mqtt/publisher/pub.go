package main

import (
	"aulas/distribuida/shared"
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

	// loop
	for i := 0; i < 10; i++ {
		// cria a mensagem
		msg := fmt.Sprintf("Mensagem %d", i)

		//publicar a mensagem
		token := client.Publish(shared.MQTTTopic, qos, false, msg)
		token.Wait()
		if token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
		fmt.Printf("Mensagem Publicada: %s\n", msg)
		time.Sleep(time.Second)
	}
}
