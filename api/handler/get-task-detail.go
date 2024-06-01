package handler

import (
	"api/gen"
	"context"
	"net/http"
)

// タスク詳細を取得する
// (GET /tasks/{taskId})
func (s *Server) GetTaskDetail(ctx context.Context, request gen.GetTaskDetailRequestObject) (gen.GetTaskDetailResponseObject, error) {
	t, err := s.tasksRepository.FindById(ctx, request.TaskId)
	if err != nil {
		return gen.GetTaskDetaildefaultJSONResponse{
			StatusCode: http.StatusInternalServerError,
			Body: gen.Error{
				Code:    "ERROR",
				Message: "タスクの取得に失敗しました",
			},
		}, err
	}

	return gen.GetTaskDetail200JSONResponse{
		Id:          t.Id(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      gen.TaskStatus(t.Status()),
	}, nil
}
