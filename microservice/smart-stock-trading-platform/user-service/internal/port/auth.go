package port

import "context"

type AuthService interface {
	GenerateToken(ctx context.Context, userID uint) (string, error)
	// ValidateToken(ctx context.Context, token string) (uint, error)
}
