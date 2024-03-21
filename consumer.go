package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "console-consumer-21564",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Subscribe to target topics
	err = consumer.SubscribeTopics([]string{"coordinates"}, nil)
	if err != nil {
		panic(err)
	}

	// Consume messages
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error while consuming message: %v\n", err)
			continue // Continue to the next iteration of the loop
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
