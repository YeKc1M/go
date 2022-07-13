package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	defer p.Close()
	topic := "quickstart"
	value := "quickstart value"
	channel := make(chan kafka.Event)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(value),
	}, channel)
	if err != nil {
		log.Fatalln(err)
	}
	e := <-channel
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		err = m.TopicPartition.Error
		log.Fatalln(err)
	}
	log.Println("success")
}
