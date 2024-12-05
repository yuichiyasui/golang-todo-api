package repository

import "database/sql"

type UsersRepository struct {
	db *sql.DB
}
