package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/DevNiklesh/kafka-producer-consumer/ms-producer/kafka"
)

var (
	frequency time.Duration = 10 * time.Second
	count                   = 0
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func messageLog(t time.Time) {
	connK := kafka.GetNewKafka()

	err := connK.Publish(fmt.Sprintf("New Message %d: %s", count, t.String()))
	if err != nil {
		log.Fatalf("kafka publish: %s", err)
	}

	count++
}

func StartCron() {
	doEvery(frequency, messageLog)
}
