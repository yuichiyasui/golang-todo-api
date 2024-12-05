package domain

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (*Email, error) {
	if !EmailRegex.MatchString(value) {
		return nil, errors.New("メールアドレスの形式が不正です")
	}

	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
