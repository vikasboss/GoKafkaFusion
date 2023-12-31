package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-container:9092",
		// consumer id
		"client.id": "go-first-consumer",
		// kafka group id
		"group.id": "first-consumers",
		// to consume all messages
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer:", err.Error())
	}
	topics := []string{"first"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("Error in subscribing to the topic", err.Error())
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		} else {
			log.Print("error in consuming:", err.Error())
		}
	}

}
