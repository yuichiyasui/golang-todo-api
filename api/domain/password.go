package domain

import "golang.org/x/crypto/bcrypt"

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	psw := []byte(value)
	hashed, err := bcrypt.GenerateFromPassword(psw, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Password{value: string(hashed)}, nil
}

func (p Password) Value() string {
	return p.value
}
