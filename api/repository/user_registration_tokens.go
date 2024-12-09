package repository

import (
	"api/domain"
	domainRepository "api/domain/repository"
	"api/model"
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRegistrationTokensRepository struct {
	db *sql.DB
}

func NewUserRegistrationTokensRepository(db *sql.DB) domainRepository.UserRegistrationTokensRepositoryInterface {
	return &UserRegistrationTokensRepository{
		db: db,
	}
}

func (r *UserRegistrationTokensRepository) Save(ctx context.Context, token string, email string, expiresAt time.Time) error {
	urt := model.UserRegistrationToken{
		Token:     token,
		Email:     email,
		ExpiresAt: expiresAt,
	}

	err := urt.Upsert(ctx, r.db, boil.Infer(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRegistrationTokensRepository) FindByToken(ctx context.Context, token string) (*domain.UserRegistrationToken, error) {
	urt, err := model.UserRegistrationTokens(
		model.UserRegistrationTokenWhere.Token.EQ(token),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return domain.NewUserRegistrationToken(urt.Token, &urt.ExpiresAt, urt.Email)
}

func (r *UserRegistrationTokensRepository) DeleteByToken(ctx context.Context, token string) error {
	_, err := model.UserRegistrationTokens(
		model.UserRegistrationTokenWhere.Token.EQ(token),
	).DeleteAll(ctx, r.db)
	if err != nil {
		return err
	}

	return nil
}
