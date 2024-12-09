package domain

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type Email struct {
	value string
}

var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, errors.New("メールアドレスは必須です")
	}

	length := utf8.RuneCountInString(value)
	if length > 100 {
		return nil, errors.New("メールアドレスは100文字以下で入力してください")
	}

	if !EmailRegex.MatchString(value) {
		return nil, errors.New("メールアドレスの形式が不正です")
	}

	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
