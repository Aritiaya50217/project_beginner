package queue

import (
	"context"
	"encoding/json"
	"notification-worker/internal/domain"
	"notification-worker/internal/ports"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(broker, topic, group string) ports.QueueConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			Topic:   topic,
			GroupID: group,
		}),
	}
}

func (c *KafkaConsumer) ConsumeContractEvents(handler func(event domain.ContractCreatedEvent) error) error {
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}
		var ev domain.ContractCreatedEvent
		_ = json.Unmarshal(m.Value, &ev)
		if err := handler(ev); err != nil {
			return err
		}
	}
}
