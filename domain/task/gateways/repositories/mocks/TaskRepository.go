// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/ManuelP84/calendar/domain/task/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TaskRepository) Close() {
	_m.Called()
}

// DeleteTaskById provides a mock function with given fields: ctx, id
func (_m *TaskRepository) DeleteTaskById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTasks provides a mock function with given fields: ctx
func (_m *TaskRepository) GetTasks(ctx context.Context) ([]*models.Task, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*models.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertTask provides a mock function with given fields: ctx, task
func (_m *TaskRepository) InsertTask(ctx context.Context, task *models.Task) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchTaskById provides a mock function with given fields: ctx, id
func (_m *TaskRepository) SearchTaskById(ctx context.Context, id string) (*models.Task, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Task); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskRepository(t mockConstructorTestingTNewTaskRepository) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
