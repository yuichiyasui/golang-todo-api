package domain

import (
	"os"
	"time"

	"github.com/google/uuid"
)

// ユーザー登録用のトークン
type UserRegistrationToken struct {
	value     string
	expiresAt time.Time
	email     Email
}

func NewUserRegistrationToken(token string, expiresAt *time.Time, email string) (*UserRegistrationToken, error) {
	emailDomain, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	if token != "" {
		return &UserRegistrationToken{
			value:     token,
			expiresAt: *expiresAt,
			email:     *emailDomain,
		}, nil
	}

	value, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UserRegistrationToken{
		value:     value.String(),
		expiresAt: time.Now().Add(24 * time.Hour),
		email:     *emailDomain,
	}, nil
}

func (t *UserRegistrationToken) Value() string {
	return t.value
}

func (t *UserRegistrationToken) ExpiresAt() time.Time {
	return t.expiresAt
}

func (t *UserRegistrationToken) Email() Email {
	return t.email
}

func (t *UserRegistrationToken) GenerateSignUpUrl() string {
	origin := os.Getenv("APP_URL")
	return origin + "/sign-up?token=" + t.value
}

func (t *UserRegistrationToken) IsExpired() bool {
	return t.expiresAt.Before(time.Now())
}
