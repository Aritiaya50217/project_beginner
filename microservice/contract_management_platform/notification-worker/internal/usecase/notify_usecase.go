package usecase

import (
	"notification-worker/internal/domain"
	"notification-worker/internal/ports"
)

type NotifyUsecase struct {
	consumer ports.QueueConsumer
	notifier ports.Notifier
}

func NewNotifyUsecase(consumer ports.QueueConsumer, notifier ports.Notifier) *NotifyUsecase {
	return &NotifyUsecase{consumer: consumer, notifier: notifier}
}

func (usecase *NotifyUsecase) Start() error {
	return usecase.consumer.ConsumeContractEvents(func(event domain.ContractCreatedEvent) error {
		return usecase.notifier.Send(event.UserID, event.Message)
	})
}
