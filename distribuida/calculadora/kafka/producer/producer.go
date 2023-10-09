package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	// definir broker kafka & tópico
	broker := "localhost:9092"
	topic := "my-topic"

	// configurar o produto
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %s", err)
	}
	defer producer.Close()

	// definir canal para produzir mensagens
	deliveryChan := make(chan kafka.Event)

	// envio das mensagens
	for i := 0; i < 5; i++ {
		// criar mensagem
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("Message %d", i)),
		}

		// colocar mensagem na fila (assíncrono)
		err := producer.Produce(message, deliveryChan)
		if err != nil {
			log.Printf("Failed to produce message: %s", err)
		} else {
			e := <-deliveryChan
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v", m.TopicPartition.Error)
			} else {
				log.Printf("Message %d sent to topic %s [partition %d] at offset %v",
					i, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		}
	}
	close(deliveryChan)
}
