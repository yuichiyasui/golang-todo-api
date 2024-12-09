package repository

import (
	"api/domain"
	"api/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	usr, err := model.Users(
		model.UserWhere.Email.EQ(email),
	).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return domain.NewUser(
		fmt.Sprint(usr.ID),
		usr.Username,
		usr.Email,
		usr.Password,
	)
}

func (r *UsersRepository) Save(ctx context.Context, input domain.User) error {
	email := input.Email()
	usr := model.User{
		Username: input.Name(),
		Email:    email.Value(),
		Password: input.Password(),
	}

	err := usr.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
