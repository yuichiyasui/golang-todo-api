package handler

import (
	"api/domain/task"
	"api/gen"
	"context"
	"fmt"
	"net/http"
)

func (s *Server) UpdateTask(ctx context.Context, request gen.UpdateTaskRequestObject) (gen.UpdateTaskResponseObject, error) {
	sts, err := convertDomainStatus(request.Body.Status)
	if err != nil {
		return gen.UpdateTaskdefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	input, err := task.New(
		request.TaskId,
		request.Body.Title,
		*request.Body.Description,
		sts,
	)
	if err != nil {
		return gen.UpdateTaskdefaultJSONResponse{
			StatusCode: http.StatusBadRequest,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	updatedTask, err := s.tasksRepository.Save(ctx, *input)
	if err != nil {
		return gen.UpdateTaskdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	res := gen.UpdateTask200JSONResponse{
		Id:          updatedTask.Id(),
		Title:       updatedTask.Title(),
		Description: updatedTask.Description(),
		Status:      gen.TaskStatus(updatedTask.Status()),
	}

	return res, nil
}

func convertDomainStatus(status gen.TaskStatus) (task.TaskStatus, error) {
	switch status {
	case gen.Todo:
		return task.TaskStatusTodo, nil
	case gen.InProgress:
		return task.TaskStatusInProgress, nil
	case gen.Done:
		return task.TaskStatusDone, nil
	default:
		return "", fmt.Errorf("不正なstatusです")
	}
}
