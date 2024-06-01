package handler

import (
	"api/gen"
	"api/repository"
	"database/sql"
)

type Server struct {
	db              *sql.DB
	tasksRepository repository.TasksRepositoryInterface
}

// Make sure we conform to StrictServerInterface

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer(db *sql.DB) (*Server, error) {
	tasksRepo, err := repository.NewTasksRepository(db)
	if err != nil {
		return nil, err
	}

	return &Server{
		db:              db,
		tasksRepository: tasksRepo,
	}, nil
}
