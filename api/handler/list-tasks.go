package handler

import (
	"api/gen"
	"api/repository"
	"context"
	"strconv"
)

func (s *Server) ListTasks(ctx context.Context, request gen.ListTasksRequestObject) (gen.ListTasksResponseObject, error) {
	tasks, err := repository.GetTasks(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	res := gen.ListTasks200JSONResponse{}
	for _, v := range tasks {
		res = append(res, gen.Task{
			Id:          strconv.FormatUint(v.ID, 10),
			Title:       v.Title,
			Description: v.Description.String,
			Status:      gen.TaskStatus(v.Status),
		})
	}

	return gen.ListTasks200JSONResponse(res), nil
}
