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

	// Sets your Google Cloud Platform project ID.
	projectID := "pcd40813"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)

}
