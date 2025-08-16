package usecase

import (
	"context"
	"smart-stock-trading-platform-order-service/internal/domain"
	"smart-stock-trading-platform-order-service/internal/port"
	"smart-stock-trading-platform-order-service/internal/utils"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrder(ctx context.Context, id uint) (*domain.Order, error)
	GetOrdersByUser(ctx context.Context, userID uint, offset, limit int64) ([]*domain.Order, int64, error)
	DeleteOrder(ctx context.Context, id uint) error
}

type orderUsecase struct {
	repo port.OrderRepository
}

func NewOrderUsecase(repo port.OrderRepository) OrderUsecase {
	return &orderUsecase{repo: repo}
}

func (u *orderUsecase) CreateOrder(ctx context.Context, order *domain.Order) error {
	if err := u.repo.Create(order); err != nil {
		utils.ErrorLogger.Printf("Failed to create order: %+v, error: %v", order, err)
		return err
	}
	return nil
}

func (u *orderUsecase) GetOrder(ctx context.Context, id uint) (*domain.Order, error) {
	return u.repo.FindByID(id)
}

func (u *orderUsecase) GetOrdersByUser(ctx context.Context, userID uint, offset, limit int64) ([]*domain.Order, int64, error) {
	return u.repo.FindByUserID(userID, offset, limit)
}

func (u *orderUsecase) DeleteOrder(ctx context.Context, id uint) error {
	return u.repo.DeleteOrder(ctx, id)
}
