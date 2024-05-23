package repository

import (
	"api/domain"
	"api/model"
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TasksRepositoryInterface interface {
	GetTasks(ctx context.Context) (model.TaskSlice, error)
	CreateTask(ctx context.Context, task domain.Task) (*model.Task, error)
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

func (r *TasksRepository) CreateTask(ctx context.Context, task domain.Task) (*model.Task, error) {
	var status model.TasksStatus
	switch task.Status {
	case domain.TaskStatusInProgress:
		status = model.TasksStatusIN_PROGRESS
	case domain.TaskStatusDone:
		status = model.TasksStatusDONE
	default:
		status = model.TasksStatusTODO
	}

	t := model.Task{
		Title:       task.Title,
		Description: null.StringFrom(task.Description),
		Status:      status,
	}
	err := t.Insert(ctx, r.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	return &t, nil
}
