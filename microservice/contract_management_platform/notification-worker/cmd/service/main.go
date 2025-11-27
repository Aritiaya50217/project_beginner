package main

import (
	"log"
	"notification-worker/internal/adapters/notifier"
	"notification-worker/internal/adapters/queue"
	"notification-worker/internal/usecase"
	"os"
)

func main() {
	consumer := queue.NewKafkaConsumer(
		os.Getenv("KAFKA_BROKER"),
		"contract.events",
		"notification-worker",
	)
	notifier := notifier.NewLogNotifier()
	uc := usecase.NewNotifyUsecase(consumer, notifier)

	log.Println("notification-worker started...")
	if err := uc.Start(); err != nil {
		log.Fatal(err)
	}
}
