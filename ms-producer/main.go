package main

import (
	"github.com/DevNiklesh/kafka-producer-consumer/ms-producer/app"
	"github.com/DevNiklesh/kafka-producer-consumer/ms-producer/cron"
)

func main() {
	// cron job to push a message for every 20 seconds
	go cron.StartCron()

	// Start the producer app
	app.StartApplication()
}
