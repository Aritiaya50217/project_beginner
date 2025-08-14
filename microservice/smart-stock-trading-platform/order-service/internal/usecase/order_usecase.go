package usecase

import (
	"context"
	"smart-stock-trading-platform-order-service/internal/domain"
	"smart-stock-trading-platform-order-service/internal/port"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrder(ctx context.Context, id uint) (*domain.Order, error)
	GetOrdersByUser(ctx context.Context, userID uint) ([]*domain.Order, error)
}

type orderUsecase struct {
	repo      port.OrderRepository
	publisher port.EventPublisher
}

func NewOrderUsecase(repo port.OrderRepository, publisher port.EventPublisher) OrderUsecase {
	return &orderUsecase{repo: repo, publisher: publisher}
}

func (u *orderUsecase) CreateOrder(ctx context.Context, order *domain.Order) error {
	order.Status = domain.StatusPending
	if err := u.repo.Create(order); err != nil {
		return err
	}
	// publisher
	if u.publisher != nil {
		return u.publisher.PublishOrderCreated(order)
	}
	return nil
}

func (u *orderUsecase) GetOrder(ctx context.Context, id uint) (*domain.Order, error) {
	return u.repo.FindByID(id)
}

func (u *orderUsecase) GetOrdersByUser(ctx context.Context, userID uint) ([]*domain.Order, error) {
	return u.repo.FindByUserID(userID)
}
