package repository

import (
	"api/domain"
	"context"
	"time"
)

type UserRegistrationTokensRepositoryInterface interface {
	Save(ctx context.Context, token string, email string, expiresAt time.Time) error
	FindByToken(ctx context.Context, token string) (*domain.UserRegistrationToken, error)
	DeleteByToken(ctx context.Context, token string) error
}
