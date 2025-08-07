package repository

import (
	"booking-system-booking-service/internal/domain"
	"context"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *domain.Booking) error
	GetByID(ctx context.Context, id int) (*domain.Booking, error)
	GetAllBooking(ctx context.Context, offset, limit int) ([]*domain.Booking, int64, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	return r.db.WithContext(ctx).Create(booking).Error
}

func (r *bookingRepository) GetByID(ctx context.Context, id int) (*domain.Booking, error) {
	var booking domain.Booking
	if err := r.db.WithContext(ctx).Preload("Item").First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetAllBooking(ctx context.Context, offset, limit int) ([]*domain.Booking, int64, error) {
	var bookings []*domain.Booking
	var total int64

	// นับจำนวนทั้งหมด
	if err := r.db.WithContext(ctx).Model(&domain.Booking{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.WithContext(ctx).
		Preload("Item").
		Offset(offset).
		Limit(limit).
		Find(&bookings).Error

	if err != nil {
		return nil, 0, err
	}

	return bookings, total, nil
}
