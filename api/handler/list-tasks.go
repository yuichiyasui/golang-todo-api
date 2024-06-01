package handler

import (
	"api/gen"
	"context"
	"net/http"
)

func (s *Server) ListTasks(ctx context.Context, request gen.ListTasksRequestObject) (gen.ListTasksResponseObject, error) {
	tasks, err := s.tasksRepository.GetTasks(ctx)
	if err != nil {
		return gen.ListTasksdefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "",
				Message: err.Error(),
			},
		}, err
	}

	res := gen.ListTasks200JSONResponse{}
	for _, v := range tasks {
		res = append(res, gen.Task{
			Id:          v.Id(),
			Title:       v.Title(),
			Description: v.Description(),
			Status:      gen.TaskStatus(v.Status()),
		})
	}

	return gen.ListTasks200JSONResponse(res), nil
}
