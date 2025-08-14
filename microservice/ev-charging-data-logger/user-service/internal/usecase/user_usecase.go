package usecase

import (
	"context"
	"errors"
	"ev-charging-data-logger-user-service/internal/domain"
	"ev-charging-data-logger-user-service/internal/repository"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email, password string) (string, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
}

type userUsecase struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func NewUserUsecase(repo repository.UserRepository, jwtSecret []byte) UserUsecase {
	return &userUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.User) error {
	existingUser, _ := u.repo.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	// check format
	if !govalidator.IsEmail(user.Email) {
		return errors.New("invalid email format")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(u.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}
