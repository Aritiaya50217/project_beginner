package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/ports"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo ports.UserRepository
}

func NewAuthUsecase(repo ports.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) Register(email, password string) error {
	// Check existing user
	if _, err := u.repo.GetByEmail(email); err != nil {
		return fmt.Errorf("email already exists")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:    email,
		Password: string(hashed),
	}
	return u.repo.Create(user)
}
