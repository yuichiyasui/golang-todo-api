package handler

import (
	"api/gen"
	"context"
)

func (s *Server) ListTasks(ctx context.Context, request gen.ListTasksRequestObject) (gen.ListTasksResponseObject, error) {
	tasks := []gen.Task{}

	return gen.ListTasks200JSONResponse(tasks), nil
}
