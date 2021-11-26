package kafka

import "log"

type IKafka interface {
	Publish(string) error
}

type kafkaStruct struct{}

func GetNewKafka() IKafka {
	return &kafkaStruct{}
}

func (k *kafkaStruct) Publish(val string) error {
	_, err := producer.Write([]byte(val))
	if err != nil {
		log.Fatalf("producer write: %s", err)
	}

	log.Printf("Data sent: %s \n", val)
	return nil
}
