package repository

import (
	"api/domain/task"
	"api/model"
	"context"
	"database/sql"
	"strconv"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=./tasks.go -destination=./tasks_mock.go -package=repository
type TasksRepositoryInterface interface {
	GetTasks(ctx context.Context) ([]*task.Task, error)
	FindById(ctx context.Context, taskId string) (*task.Task, error)
	CreateTask(ctx context.Context, input task.Task) (*task.Task, error)
	UpdateTask(ctx context.Context, input task.Task) (*task.Task, error)
}

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) (TasksRepositoryInterface, error) {
	return &TasksRepository{db: db}, nil
}

func (r *TasksRepository) GetTasks(ctx context.Context) ([]*task.Task, error) {
	data, err := model.Tasks().All(ctx, r.db)

	if err != nil {
		return nil, err
	}

	tasks := []*task.Task{}
	for _, t := range data {
		d, err := task.New(
			strconv.FormatUint(t.ID, 10),
			t.Title,
			t.Description.String,
			task.TaskStatus(t.Status),
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &d)
	}

	return tasks, nil
}

func (r *TasksRepository) FindById(ctx context.Context, taskId string) (*task.Task, error) {
	id, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		return nil, err
	}

	m, err := model.FindTask(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	t, err := task.New(
		strconv.FormatUint(m.ID, 10),
		m.Title,
		m.Description.String,
		task.TaskStatus(m.Status.String()),
	)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *TasksRepository) CreateTask(ctx context.Context, input task.Task) (*task.Task, error) {
	var status model.TasksStatus
	switch input.Status() {
	case task.TaskStatusTodo:
		status = model.TasksStatusTODO
	case task.TaskStatusInProgress:
		status = model.TasksStatusIN_PROGRESS
	case task.TaskStatusDone:
		status = model.TasksStatusDONE
	default:
		status = model.TasksStatusTODO
	}

	t := model.Task{
		Title:       input.Title(),
		Description: null.StringFrom(input.Description()),
		Status:      status,
	}
	err := t.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	createdTask, err := task.New(
		strconv.FormatUint(t.ID, 10),
		t.Title,
		t.Description.String,
		task.TaskStatus(t.Status.String()),
	)
	if err != nil {
		return nil, err
	}

	return &createdTask, nil
}

func (r *TasksRepository) UpdateTask(ctx context.Context, input task.Task) (*task.Task, error) {
	taskId, err := strconv.ParseUint(input.Id(), 10, 64)
	if err != nil {
		return nil, err
	}

	m, err := model.FindTask(ctx, r.db, taskId)
	if err != nil {
		return nil, err
	}

	m.Title = input.Title()
	m.Description = null.StringFrom(input.Description())
	m.Status = convertStatus(input.Status())

	_, err = m.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	m, err = model.FindTask(ctx, r.db, taskId)
	if err != nil {
		return nil, err
	}

	updatedTask, err := task.New(
		strconv.FormatUint(m.ID, 10),
		m.Title,
		m.Description.String,
		task.TaskStatus(m.Status.String()),
	)
	if err != nil {
		return nil, err
	}

	return &updatedTask, nil
}

func convertStatus(status task.TaskStatus) model.TasksStatus {
	switch status {
	case task.TaskStatusInProgress:
		return model.TasksStatusIN_PROGRESS
	case task.TaskStatusDone:
		return model.TasksStatusDONE
	default:
		return model.TasksStatusTODO
	}
}
