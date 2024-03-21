package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	topic := "coordinates"

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d", i)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
		}, nil)
		if err != nil {
			fmt.Printf("Failed to produce message %d: %v\n", i, err)
		} else {
			fmt.Printf("Produced message %d: %s\n", i, value)
		}
	}

	producer.Flush(15 * 1000)
	producer.Close()
}
