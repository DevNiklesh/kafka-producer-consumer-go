package kafka

import (
	"context"
	"log"
	"strconv"

	"github.com/DevNiklesh/kafka-producer-consumer/ms-consumer/config"
	"github.com/segmentio/kafka-go"
)

var (
	conf = config.New()
)

func Consumer() {
	kafkaURL := conf.KafkaURL
	topic := conf.Topic
	partition, err := strconv.Atoi(conf.Partition)
	if err != nil {
		log.Fatalf("partition string conversion: %s", err)
	}

	connK, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
	if err != nil {
		log.Fatalf("kafka dail leader: %s", err)
	}
	defer func() {
		err := connK.Close()
		if err != nil {
			log.Fatalf("kafka connection close: %s", err)
		}
	}()

	for {
		msg, err := connK.ReadMessage(1)
		if err != nil {
			log.Fatalf("read message: %s", err)
		}

		log.Println("Received from kafka: ", string(msg.Value))
	}
}
