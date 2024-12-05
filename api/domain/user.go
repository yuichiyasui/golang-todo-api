package domain

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	id    string
	name  string
	email string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}

	length := utf8.RuneCountInString(name)
	if length > 30 {
		return fmt.Errorf("name is too long")
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}

	length := utf8.RuneCountInString(email)
	if length > 100 {
		return fmt.Errorf("email is too long")
	}

	return nil
}

func NewUser(id, name, email string) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	if err := validateEmail(email); err != nil {
		return nil, err
	}

	return &User{id: id, name: name, email: email}, nil
}
