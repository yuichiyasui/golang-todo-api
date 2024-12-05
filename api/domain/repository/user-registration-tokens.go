package repository

import "context"

type UserRegistrationTokensRepositoryInterface interface {
	Save(ctx context.Context, token string) error
}
