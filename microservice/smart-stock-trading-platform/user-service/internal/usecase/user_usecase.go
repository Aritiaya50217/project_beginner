package usecase

import (
	"context"
	"errors"
	"smart-stock-trading-platform-user-service/internal/domain"
	"smart-stock-trading-platform-user-service/internal/port"
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

	hashPassword, err := u.auth.HashPassword(ctx, password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:     email,
		Password:  string(hashPassword),
		FirstName: firstname,
		LastName:  lastname,
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return "", errors.New("invalid email")
	}

	if !u.auth.CheckPasswordHash(ctx, password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	return u.auth.GenerateToken(ctx, user.ID, user.Email)
}

func (u *UserUsecase) Create(ctx context.Context, user *domain.User) error {
	return u.repo.Create(ctx, user)
}
func (u *UserUsecase) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return u.repo.GetUserByID(ctx, id)
}
