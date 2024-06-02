package task

import (
	"fmt"
	"unicode/utf8"
)

type TaskStatus string

type Task struct {
	id          string
	title       string
	description string
	status      TaskStatus
}

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

func (t *Task) Id() string {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() TaskStatus {
	return t.status
}

func validateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("title is required")
	}
	length := utf8.RuneCountInString(title)
	if length > 30 {
		return fmt.Errorf("title is too long")
	}

	return nil
}

func validateDescription(description string) error {
	length := utf8.RuneCountInString(description)
	if length > 500 {
		return fmt.Errorf("description is too long")
	}

	return nil
}

func convertStatus(status TaskStatus) TaskStatus {
	switch status {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return status
	default:
		return TaskStatusTodo
	}
}

func New(id string, title string, description string, status TaskStatus) (*Task, error) {
	if err := validateTitle(title); err != nil {
		return nil, err
	}
	if err := validateDescription(description); err != nil {
		return nil, err
	}

	return &Task{
		id:          id,
		title:       title,
		description: description,
		status:      convertStatus(status),
	}, nil
}
