package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	topicID := "topic3"

	// Configura projeto
	projectID := "pcd40813"

	// Cria cliente
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// publica mensagem
	topic := client.Topic(topicID)
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("calculadora world"),
	})

	if res == nil {
		log.Fatalf("Failed to publish in the topic: %v", err)
	}
	defer topic.Stop()
}
