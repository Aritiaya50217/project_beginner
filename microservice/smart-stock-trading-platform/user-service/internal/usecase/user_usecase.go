package usecase

import (
	"context"
	"errors"
	"smart-stock-trading-platform-user-service/internal/domain"
	"smart-stock-trading-platform-user-service/internal/port"
	"time"
)

type UserUsecase struct {
	repo           port.UserRepository
	auth           port.AuthService
	repoForRefresh port.UserRefreshTokenRepository
}

func NewUserUsecase(repo port.UserRepository, auth port.AuthService, repoForRefresh port.UserRefreshTokenRepository) *UserUsecase {
	return &UserUsecase{repo: repo, auth: auth, repoForRefresh: repoForRefresh}
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

func (u *UserUsecase) Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return "", "", errors.New("invalid email")
	}

	if !u.auth.CheckPasswordHash(ctx, password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err = u.auth.GenerateToken(ctx, user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = u.auth.GenerateRefreshToken(ctx, user.ID)
	if err != nil {
		return "", "", err
	}

	// บันทึก refresh token ลง DB (อายุ 1 วัน)
	expireAt := time.Now().Add(24 * time.Hour).Unix()
	err = u.repoForRefresh.SaveRefreshToken(ctx, user.ID, refreshToken, expireAt)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
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

func (u *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	userID, err := u.auth.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", nil
	}

	valid, err := u.repoForRefresh.IsRefreshTokenValid(ctx, userID, refreshToken)
	if err != nil {
		return "", err
	}

	if !valid {
		return "", errors.New("refresh token is invalid or revoked")
	}

	// สร้าง access token
	return u.auth.GenerateRefreshToken(ctx, userID)
}

func (u *UserUsecase) RevokeRefreshToken(ctx context.Context, userID uint, refreshToken string) error {
	return u.repoForRefresh.RevokeRefreshToken(ctx, userID, refreshToken)
}
