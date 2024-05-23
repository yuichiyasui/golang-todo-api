package domain

type TaskStatus string

type Task struct {
	Id          string
	Title       string
	Description string
	Status      TaskStatus
}

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusDone       TaskStatus = "DONE"
)

func New(id string, title string, description string, status TaskStatus) (Task, error) {
	var sts TaskStatus
	switch status {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		sts = status
	default:
		sts = TaskStatusTodo
	}

	return Task{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      sts,
	}, nil
}
