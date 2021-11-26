package kafka

import (
	"context"
	"log"
	"strconv"

	"github.com/DevNiklesh/kafka-producer-consumer/ms-producer/config"
	"github.com/segmentio/kafka-go"
)

var (
	producer *kafka.Conn
	conf     = config.New()
)

func init() {
	kafkaURL := conf.KafkaURL
	topic := conf.Topic
	partition, err := strconv.Atoi(conf.Partition)
	if err != nil {
		log.Fatalf("partition conversion: %s", err)
	}

	// Kafka connection
	producer, err = kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
	if err != nil {
		log.Fatalf("kafka dialleader: %s", err)
	}
}
