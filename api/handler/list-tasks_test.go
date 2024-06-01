package handler_test

import (
	"api/domain/task"
	"api/repository"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestListTasks(t *testing.T) {
	t.Run("タスク一覧とstatus200を返す", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := repository.NewMockTasksRepositoryInterface(mockCtrl)
		mockCtx, _ := gin.CreateTestContext(httptest.NewRecorder())

		task1, _ := task.New(
			"1",
			"タスク名",
			"タスク説明",
			task.TaskStatusTodo,
		)
		mock.EXPECT().GetTasks(mockCtx).Return(
			task1,
		)
	})

}
