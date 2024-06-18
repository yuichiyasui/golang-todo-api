package handler

import (
	"api/domain/task"
	"api/gen"
	"api/repository"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServer_GetTaskDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		mockFunc   func(mock *repository.MockTasksRepositoryInterface)
		request    gen.GetTaskDetailRequestObject
		response   string
		httpStatus int
		wantErr    bool
	}{
		{
			name: "タスク詳細を取得する",
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				obj, _ := task.New("1", "title", "description", task.TaskStatusTodo)
				m.EXPECT().FindById(gomock.Any(), "1").Return(obj, nil)
			},
			request: gen.GetTaskDetailRequestObject{
				TaskId: "1",
			},
			response: `
				{
					"id"			:	"1",
					"title"			:	"title",
					"description"	:	"description",
					"status"		:	"TODO"
				}
			`,
			httpStatus: 200,
			wantErr:    false,
		},
		{
			name: "タスク詳細の取得に失敗した場合、ステータス500を返す",
			mockFunc: func(m *repository.MockTasksRepositoryInterface) {
				m.EXPECT().FindById(gomock.Any(), "1").Return(nil, assert.AnError)
			},
			request: gen.GetTaskDetailRequestObject{
				TaskId: "1",
			},
			response: `
				{
					"code"		:	"ERROR",
					"message"	:	"タスクの取得に失敗しました"
				}
			`,
			httpStatus: 500,
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

			res, err := s.GetTaskDetail(mockCtx, tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			err = res.VisitGetTaskDetailResponse(recorder)
			assert.NoError(t, err)

			assert.Equal(t, tt.httpStatus, recorder.Code)
			assert.JSONEq(t, tt.response, recorder.Body.String())
		})
	}
}
