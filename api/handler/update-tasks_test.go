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

func TestUpdateTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	desc := "タスク説明"

	tests := []struct {
		name       string
		mockFunc   func(mock *repository.MockTasksRepositoryInterface)
		request    gen.UpdateTaskRequestObject
		response   string
		httpStatus int
		wantErr    bool
	}{
		{
			name: "タスクを更新して更新後のタスクを返す",
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				obj, _ := task.New("1", "title", "タスク説明", task.TaskStatusInProgress)
				m.EXPECT().UpdateTask(gomock.Any(), *obj).Return(obj, nil)
			},
			request: gen.UpdateTaskRequestObject{
				TaskId: "1",
				Body: &gen.UpdateTaskJSONRequestBody{
					Title:       "title",
					Description: &desc,
					Status:      gen.InProgress,
				},
			},
			response: `
				{
					"id"			:	"1",
					"title"			:	"title",
					"description"	:	"タスク説明",
					"status"		:	"IN_PROGRESS"
				}
			`,
			httpStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "タスクのステータスの変換に失敗した場合、ステータス400を返す",
			request: gen.UpdateTaskRequestObject{
				TaskId: "1",
				Body: &gen.UpdateTaskJSONRequestBody{
					Title:       "title",
					Description: &desc,
					Status:      "invalid",
				},
			},
			response:   `{"code":"","message":"不正なstatusです"}`,
			httpStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name: "ドメインオブジェクトの生成に失敗した場合、ステータス400を返す",
			request: gen.UpdateTaskRequestObject{
				TaskId: "1",
				Body: &gen.UpdateTaskJSONRequestBody{
					Title:       "",
					Description: &desc,
					Status:      gen.InProgress,
				},
			},
			response:   `{"code":"","message":"title is required"}`,
			httpStatus: http.StatusBadRequest,
			wantErr:    true,
		},
		{
			name: "タスクの更新に失敗した場合、ステータス500を返す",
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				m.EXPECT().UpdateTask(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error message"))
			},
			request: gen.UpdateTaskRequestObject{
				TaskId: "1",
				Body: &gen.UpdateTaskJSONRequestBody{
					Title:       "title",
					Description: &desc,
					Status:      gen.InProgress,
				},
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

			res, err := s.UpdateTask(mockCtx, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			err = res.VisitUpdateTaskResponse(recorder)
			assert.NoError(t, err)

			assert.Equal(t, tt.httpStatus, recorder.Code)
			assert.JSONEq(t, tt.response, recorder.Body.String())
		})
	}
}
