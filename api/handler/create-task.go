package handler

import (
	"api/domain/task"
	"api/gen"
	"context"
	"net/http"
)

func (s *Server) CreateTask(ctx context.Context, request gen.CreateTaskRequestObject) (gen.CreateTaskResponseObject, error) {
	input, err := task.New(
		"",
		request.Body.Title,
		*request.Body.Description,
		"",
	)
	if err != nil {
		return gen.CreateTaskdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	createdTask, err := s.tasksRepository.CreateTask(ctx, *input)
	if err != nil {
		return gen.CreateTaskdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	return gen.CreateTask200JSONResponse{
		Id: createdTask.Id(),
	}, nil
}
