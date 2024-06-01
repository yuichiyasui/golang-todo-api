package handler

import (
	"api/domain"
	"api/gen"
	"context"
	"fmt"
	"strconv"
)

func (s *Server) UpdateTask(ctx context.Context, request gen.UpdateTaskRequestObject) (gen.UpdateTaskResponseObject, error) {
	sts, err := convertDomainStatus(*request.Body.Status)
	if err != nil {
		return nil, err
	}

	task, err := domain.New(
		request.TaskId,
		request.Body.Title,
		*request.Body.Description,
		sts,
	)
	if err != nil {
		return nil, err
	}

	updatedTask, err := s.tasksRepository.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	res := gen.UpdateTask200JSONResponse{
		Id:          strconv.FormatUint(updatedTask.ID, 10),
		Title:       updatedTask.Title,
		Description: updatedTask.Description.String,
		Status:      gen.TaskStatus(updatedTask.Status),
	}

	return res, nil
}

func convertDomainStatus(status gen.TaskStatus) (domain.TaskStatus, error) {
	switch status {
	case gen.Todo:
		return domain.TaskStatusTodo, nil
	case gen.InProgress:
		return domain.TaskStatusInProgress, nil
	case gen.Done:
		return domain.TaskStatusDone, nil
	default:
		return "", fmt.Errorf("不正なstatusです")
	}
}
