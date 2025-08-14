package publisher

import (
	"context"
	"encoding/json"
	"log"
	"smart-stock-trading-platform-order-service/internal/domain"
	"smart-stock-trading-platform-order-service/internal/port"
	"smart-stock-trading-platform-order-service/internal/utils"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type kafkaPublisher struct {
	writer *kafka.Writer
}

// NewKafkaPublisher สร้าง publisher พร้อม Kafka writer
func NewKafkaPublisher(broker, topic string) port.EventPublisher {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &kafkaPublisher{writer: writer}
}

func (k *kafkaPublisher) PublishOrderCreated(order *domain.Order) error {
	msgBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	key := []byte(strconv.Itoa(int(order.ID))) // แปลง ID เป็น string ก่อนเป็น []byte

	// Retry 3 ครั้ง เผื่อ broker ยังไม่ ready
	for i := 0; i < 3; i++ {
		err = k.writer.WriteMessages(context.Background(), kafka.Message{
			Key:   key,
			Value: msgBytes,
		})
		if err == nil {
			break
		}
		utils.Logger.Warn("Publish failed, retrying...", zap.Error(err))
		time.Sleep(time.Second * 2)
	}

	if err != nil {
		utils.Logger.Error("Failed to publish order created", zap.Error(err))
		return err
	}

	utils.Logger.Info("Published order created", zap.Int("orderID", int(order.ID)))
	return nil
}

func (k *kafkaPublisher) PublishOrderUpdated(order *domain.Order) error {
	msgBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = k.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(string(order.ID)),
			Value: msgBytes,
		},
	)
	if err != nil {
		log.Println("Failed to publish order updated:", err)
		return err
	}

	log.Println("Published order updated:", order.ID)
	return nil
}
