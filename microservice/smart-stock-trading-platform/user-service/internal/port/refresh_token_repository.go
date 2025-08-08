package port

import "context"

type UserRefreshTokenRepository interface {
	SaveRefreshToken(ctx context.Context, userID uint, refreshToken string, expiresAt int64) error
	IsRefreshTokenValid(ctx context.Context, userID uint, refreshToken string) (bool, error)
	RevokeRefreshToken(ctx context.Context, userID uint, refreshToken string) error
}
