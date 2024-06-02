package handler

import (
	"api/domain/task"
	"api/gen"
	"api/repository"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("タスク一覧とstatus200を返す", func(t *testing.T) {
		task1, _ := task.New(
			"1",
			"タスク名",
			"タスク説明",
			task.TaskStatusTodo,
		)
		task2, _ := task.New(
			"2",
			"タスク名2",
			"タスク説明",
			task.TaskStatusInProgress,
		)
		mockTasks := []*task.Task{
			task1,
			task2,
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		recorder := httptest.NewRecorder()

		mock := repository.NewMockTasksRepositoryInterface(mockCtrl)
		mock.EXPECT().GetTasks(gomock.Any()).Return(mockTasks, nil)

		server := &Server{
			db:              &sql.DB{},
			tasksRepository: mock,
		}
		mockCtx, _ := gin.CreateTestContext(recorder)
		res, err := server.ListTasks(mockCtx, gen.ListTasksRequestObject{})

		assert.NoError(t, err)

		err = res.VisitListTasksResponse(recorder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, `
			[
				{
					"id"			:	"1", 
					"title"			:	"タスク名",
					"description"	:	"タスク説明", 
					"status"		:	"TODO"
				},
				{
					"id"			:	"2", 
					"title"			:	"タスク名2",
					"description"	:	"タスク説明", 
					"status"		:	"IN_PROGRESS"
				}
			]
			`,
			recorder.Body.String(),
		)
	})

	t.Run("リポジトリでタスク一覧の取得時にエラーが発生した場合、status500を返す", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		recorder := httptest.NewRecorder()

		mock := repository.NewMockTasksRepositoryInterface(mockCtrl)
		mock.EXPECT().GetTasks(gomock.Any()).Return(nil, fmt.Errorf("error"))

		server := &Server{
			db:              &sql.DB{},
			tasksRepository: mock,
		}
		mockCtx, _ := gin.CreateTestContext(recorder)
		res, err := server.ListTasks(mockCtx, gen.ListTasksRequestObject{})
		assert.Error(t, err)

		err = res.VisitListTasksResponse(recorder)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.JSONEq(t, `
			{
				"code"		:"ERROR",
				"message"	:"error"
			}
		`, recorder.Body.String())
	})
}
