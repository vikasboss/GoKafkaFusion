package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "my_Group_set",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"my_topic_set", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
