package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/pkg/utils"
	"auth-service/internal/ports"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo ports.UserRepository
}

func NewUserUsecase(repo ports.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// check password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (u *UserUsecase) GetProfile(userID uint) (*domain.User, error) {
	return u.repo.GetByID(userID)
}
