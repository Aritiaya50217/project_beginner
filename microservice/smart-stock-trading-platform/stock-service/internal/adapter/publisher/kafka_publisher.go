package publisher

import (
	"fmt"
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type kafkaPublisher struct {
	brokers []string
}

func NewKafkaPublisher(brokers []string) *kafkaPublisher {
	return &kafkaPublisher{brokers: brokers}
}

func (k *kafkaPublisher) PublishPriceUpdate(stock domain.Stock) error {
	// Mock: แค่ print log
	fmt.Printf("Publishing to Kafka: %+v\n", stock)
	return nil
}
