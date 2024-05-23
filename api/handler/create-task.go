package handler

import (
	"api/domain"
	"api/gen"
	"context"
	"net/http"
	"strconv"
)

func (s *Server) CreateTask(ctx context.Context, request gen.CreateTaskRequestObject) (gen.CreateTaskResponseObject, error) {
	t, err := domain.New(
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

	task, err := s.tasksRepository.CreateTask(ctx, t)
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
		Id: strconv.FormatUint(task.ID, 10),
	}, nil
}
