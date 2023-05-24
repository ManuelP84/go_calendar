package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/ManuelP84/calendar/domain/task/events"
	busMocks "github.com/ManuelP84/calendar/domain/task/gateways/bus/mocks"
	repoMocks "github.com/ManuelP84/calendar/domain/task/gateways/repositories/mocks"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/stretchr/testify/assert"
)

func TestInsertTaskNil(t *testing.T) {
	ctx := context.TODO()
	var task *models.Task
	expectedError := errors.New("task can't be null")

	usecase := NewInsertTask(nil, nil)
	err := usecase.InsertTask(ctx, task)

	assert.Equal(t, expectedError, err)
}

func TestInsertTask(t *testing.T) {
	ctx := context.TODO()
	task := &models.Task{
		Id:          "1",
		Title:       "First task",
		Description: "First task description",
	}

	taskCreatedEvent := events.TaskEvent{
		EventType: events.TaskCreatedEvent,
		Task:      task,
	}

	taskRepository := repoMocks.NewTaskRepository(t)
	taskProducer := busMocks.NewTaskBus(t)

	usecase := NewInsertTask(taskRepository, taskProducer)

	taskRepository.On("InsertTask", ctx, task).Return(nil)
	taskProducer.On("Publish", taskCreatedEvent).Return(nil)

	err := usecase.InsertTask(ctx, task)

	assert.Equal(t, nil, err)
}
