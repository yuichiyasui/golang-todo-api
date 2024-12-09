package domain

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	id       string
	name     string
	email    Email
	password string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) Password() string {
	return u.password
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

func NewUser(id, name, email, password string) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	e, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		id:       id,
		name:     name,
		email:    *e,
		password: password,
	}, nil
}
