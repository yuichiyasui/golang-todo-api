package repository

import (
	"api/domain"
	"context"
)

type UsersRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, input domain.User) (*domain.User, error)
}
