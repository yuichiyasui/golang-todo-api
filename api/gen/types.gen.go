// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package gen

// Defines values for TaskStatus.
const (
	Done       TaskStatus = "done"
	InProgress TaskStatus = "inProgress"
	Todo       TaskStatus = "todo"
)

// Error defines model for Error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Task defines model for Task.
type Task struct {
	Description string     `json:"description"`
	Id          string     `json:"id"`
	Status      TaskStatus `json:"status"`
	Title       string     `json:"title"`
}

// TaskStatus defines model for TaskStatus.
type TaskStatus string

// CreateTaskJSONBody defines parameters for CreateTask.
type CreateTaskJSONBody struct {
	// Description タスクの説明
	Description *string `json:"description,omitempty"`

	// Title タスク名
	Title string `json:"title"`
}

// UpdateTaskJSONBody defines parameters for UpdateTask.
type UpdateTaskJSONBody struct {
	// Description タスクの説明
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`

	// Title タスク名
	Title string `json:"title"`
}

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody CreateTaskJSONBody

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody UpdateTaskJSONBody
