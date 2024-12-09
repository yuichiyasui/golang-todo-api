// Code generated by MockGen. DO NOT EDIT.
// Source: ./tasks.go

// Package repository is a generated GoMock package.
package repository

import (
	task "api/domain/task"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTasksRepositoryInterface is a mock of TasksRepositoryInterface interface.
type MockTasksRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTasksRepositoryInterfaceMockRecorder
}

// MockTasksRepositoryInterfaceMockRecorder is the mock recorder for MockTasksRepositoryInterface.
type MockTasksRepositoryInterfaceMockRecorder struct {
	mock *MockTasksRepositoryInterface
}

// NewMockTasksRepositoryInterface creates a new mock instance.
func NewMockTasksRepositoryInterface(ctrl *gomock.Controller) *MockTasksRepositoryInterface {
	mock := &MockTasksRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTasksRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTasksRepositoryInterface) EXPECT() *MockTasksRepositoryInterfaceMockRecorder {
	return m.recorder
}

// FindById mocks base method.
func (m *MockTasksRepositoryInterface) FindById(ctx context.Context, taskId string) (*task.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, taskId)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockTasksRepositoryInterfaceMockRecorder) FindById(ctx, taskId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockTasksRepositoryInterface)(nil).FindById), ctx, taskId)
}

// GetTasks mocks base method.
func (m *MockTasksRepositoryInterface) GetTasks(ctx context.Context) ([]*task.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", ctx)
	ret0, _ := ret[0].([]*task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockTasksRepositoryInterfaceMockRecorder) GetTasks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockTasksRepositoryInterface)(nil).GetTasks), ctx)
}

// Save mocks base method.
func (m *MockTasksRepositoryInterface) Save(ctx context.Context, input task.Task) (*task.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, input)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockTasksRepositoryInterfaceMockRecorder) Save(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTasksRepositoryInterface)(nil).Save), ctx, input)
}
