package repository

import (
	"api/model"
	"context"
	"database/sql"
)

type TasksRepositoryInterface interface {
	GetTasks(ctx context.Context) (model.TaskSlice, error)
}

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) (*TasksRepository, error) {
	return &TasksRepository{db: db}, nil
}

func (r *TasksRepository) GetTasks(ctx context.Context) (model.TaskSlice, error) {
	tasks, err := model.Tasks().All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
