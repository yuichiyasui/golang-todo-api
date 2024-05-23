package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewDBClient() (*sql.DB, error) {
	locale, _ := time.LoadLocation("Asia/Tokyo")
	c := mysql.Config{
		DBName:    "golang_todo_api",
		User:      "root",
		Passwd:    "root",
		ParseTime: true,
		Loc:       locale,
	}
	dsn := c.FormatDSN()

	return sql.Open("mysql", dsn)
}
