package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
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
		log.Fatalf("Failed to create a publisher", err)
	}

	// subscreve ao tópico / recebe mensagem
	sub := client.Subscription(topicID)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(m)
		m.Ack()
	})
	if err != context.Canceled {
		log.Fatalf("Failed to receive message", err)
	}
}
