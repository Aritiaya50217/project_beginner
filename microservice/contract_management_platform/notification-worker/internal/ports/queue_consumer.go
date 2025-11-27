package ports

import "notification-worker/internal/domain"

type QueueConsumer interface {
	ConsumeContractEvents(handler func(event domain.ContractCreatedEvent) error) error
}
