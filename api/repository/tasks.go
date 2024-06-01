package repository

import (
	"api/domain"
	"api/model"
	"context"
	"database/sql"
	"strconv"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TasksRepositoryInterface interface {
	GetTasks(ctx context.Context) (model.TaskSlice, error)
	CreateTask(ctx context.Context, task domain.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, task domain.Task) (*model.Task, error)
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

func (r *TasksRepository) UpdateTask(ctx context.Context, task domain.Task) (*model.Task, error) {
	taskId, err := strconv.ParseUint(task.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	m, err := model.FindTask(ctx, r.db, taskId)
	if err != nil {
		return nil, err
	}

	m.Title = task.Title
	m.Description = null.StringFrom(task.Description)
	m.Status = convertStatus(task.Status)

	_, err = m.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return m, nil
}

func convertStatus(status domain.TaskStatus) model.TasksStatus {
	switch status {
	case domain.TaskStatusInProgress:
		return model.TasksStatusIN_PROGRESS
	case domain.TaskStatusDone:
		return model.TasksStatusDONE
	default:
		return model.TasksStatusTODO
	}
}
