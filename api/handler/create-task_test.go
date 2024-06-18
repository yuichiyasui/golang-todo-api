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

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	desc := "タスク説明"

	tests := []struct {
		name       string
		mockFunc   func(mock *repository.MockTasksRepositoryInterface)
		request    gen.CreateTaskRequestObject
		response   string
		httpStatus int
		wantErr    bool
	}{
		{
			name: "タスクの作成に成功した場合、作成したタスクのIDとstatus200を返す",
			request: gen.CreateTaskRequestObject{
				Body: &gen.CreateTaskJSONRequestBody{
					Title:       "タスク名",
					Description: &desc,
				},
			},
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				task1, _ := task.New(
					"1",
					"タスク名",
					"タスク説明",
					task.TaskStatusTodo,
				)
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(task1, nil)
			},
			response:   `{"id":"1"}`,
			httpStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "ドメインオブジェクトの生成に失敗した場合、エラーメッセージとstatus500を返す",
			request: gen.CreateTaskRequestObject{
				Body: &gen.CreateTaskJSONRequestBody{
					Title:       "",
					Description: &desc,
				},
			},
			response:   `{"code":"","message":"title is required"}`,
			httpStatus: http.StatusInternalServerError,
			wantErr:    true,
		},
		{
			name: "リポジトリのSaveメソッドがエラーを返した場合、エラーメッセージとstatus500を返す",
			request: gen.CreateTaskRequestObject{
				Body: &gen.CreateTaskJSONRequestBody{
					Title:       "タスク名",
					Description: &desc,
				},
			},
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				m.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error message"))
			},
			response:   `{"code":"","message":"error message"}`,
			httpStatus: http.StatusInternalServerError,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := repository.NewMockTasksRepositoryInterface(mockCtrl)
			if tt.mockFunc != nil {
				tt.mockFunc(mockRepo)
			}

			s := &Server{
				db:              &sql.DB{},
				tasksRepository: mockRepo,
			}

			recorder := httptest.NewRecorder()
			mockCtx, _ := gin.CreateTestContext(recorder)

			res, err := s.CreateTask(mockCtx, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			err = res.VisitCreateTaskResponse(recorder)
			assert.NoError(t, err)

			assert.Equal(t, tt.httpStatus, recorder.Code)
			assert.JSONEq(t, tt.response, recorder.Body.String())
		})
	}
}
