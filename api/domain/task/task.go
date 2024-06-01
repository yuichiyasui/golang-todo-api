package task

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

func convertStatus(status TaskStatus) TaskStatus {
	switch status {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone:
		return status
	default:
		return TaskStatusTodo
	}
}

func New(id string, title string, description string, status TaskStatus) (Task, error) {
	return Task{
		id:          id,
		title:       title,
		description: description,
		status:      convertStatus(status),
	}, nil
}
