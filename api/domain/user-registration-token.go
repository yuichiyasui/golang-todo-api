package domain

import (
	"os"

	"github.com/google/uuid"
)

// ユーザー登録用のトークン
type UserRegistrationToken struct {
	value string
}

func NewUserRegistrationToken() (*UserRegistrationToken, error) {
	value, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UserRegistrationToken{value: value.String()}, nil
}

func (t *UserRegistrationToken) Value() string {
	return t.value
}

func (t *UserRegistrationToken) GenerateSignUpUrl() string {
	origin := os.Getenv("APP_URL")
	return origin + "/sign-up?token=" + t.value
}
