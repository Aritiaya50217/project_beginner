package db

import (
	"car-sure-microservice/user-service/internal/domain"
	"car-sure-microservice/user-service/internal/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Constructor Function
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
