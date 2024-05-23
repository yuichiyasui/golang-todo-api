package handler

import (
	"api/gen"
	"database/sql"
)

type Server struct {
	DB *sql.DB
}

// Make sure we conform to StrictServerInterface

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}
