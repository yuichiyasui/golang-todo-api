package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBClient() (*sql.DB, error) {
	dbname := "golang_todo_api"
	user := "root"
	password := "root"

	return sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
}
