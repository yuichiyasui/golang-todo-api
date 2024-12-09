package repository

import (
	"api/domain"
	"context"
)

type UsersRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Save(ctx context.Context, input domain.User) error
}
