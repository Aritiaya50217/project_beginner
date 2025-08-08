package usecase

import (
	"context"
	"errors"
	"smart-stock-trading-platform-user-service/internal/domain"
	"smart-stock-trading-platform-user-service/internal/port"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo port.UserRepository
	auth port.AuthService
}

func NewUserUsecase(repo port.UserRepository, auth port.AuthService) *UserUsecase {
	return &UserUsecase{repo: repo, auth: auth}
}

func (u *UserUsecase) Register(ctx context.Context, email, password, firstname, lastname string) (*domain.User, error) {
	existing, _ := u.repo.FindByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Email:     email,
		Password:  string(hash),
		FirstName: firstname,
		LastName:  lastname,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
