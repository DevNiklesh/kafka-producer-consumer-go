package config

import "os"

type Config struct {
	// Application
	PORT string

	KafkaURL  string
	Topic     string
	Partition string
}

func New() *Config {
	return &Config{
		PORT:      getEnv("PORT", ":8080"),
		KafkaURL:  getEnv("KAFKA_URL", "kafka:9092"),
		Topic:     getEnv("TOPIC", "main-topic"),
		Partition: getEnv("PARTITION", "0"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
