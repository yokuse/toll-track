package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"toll-calculator/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// interface for microservice to implement so that
// we reduce dependency on kafka
type DataProducer interface {
	PushData(types.OBUData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
	topic string
}

func NewKafkaProducer(topic string) (DataProducer, error) {
	// actual kafka client producer APIs
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatal("Failed to create Kafka producer")
		return nil, err
	}

	return &KafkaProducer{
		producer: p,
		topic: topic,
	}, err
}

func (p *KafkaProducer) PushData(data types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)

	return err
}
