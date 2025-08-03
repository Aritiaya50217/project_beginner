package application

import (
	"car-sure-microservice/user-service/internal/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("my-secret-key")

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Register(email, name, password string) error {
	if (email == "") || password == "" {
		return errors.New("email and password required.")
	}

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	existUser, err := u.repo.FindByEmail(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existUser != nil {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:    email,
		Name:     name,
		Password: string(hashed),
	}
	return u.repo.Create(user)
}
