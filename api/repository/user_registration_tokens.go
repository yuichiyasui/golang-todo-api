package repository

import (
	"api/domain/repository"
	"api/model"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRegistrationTokensRepository struct {
	db *sql.DB
}

func NewUserRegistrationTokensRepository(db *sql.DB) repository.UserRegistrationTokensRepositoryInterface {
	return &UserRegistrationTokensRepository{
		db: db,
	}
}

func (r *UserRegistrationTokensRepository) Save(ctx context.Context, token string) error {
	urt := model.UserRegistrationToken{
		Token: token,
	}
	err := urt.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}