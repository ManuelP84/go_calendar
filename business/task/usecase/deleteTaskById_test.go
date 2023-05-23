package usecase

import (
	"context"
	"testing"

	repoMocks "github.com/ManuelP84/calendar/domain/task/gateways/repositories/mocks"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTaskById(t *testing.T) {
	ctx := context.TODO()
	task := &models.Task{
		Id:          "1",
		Title:       "First task",
		Description: "First task description",
	}

	taskRepository := repoMocks.NewTaskRepository(t)

	usecase := NewDeleteTaskById(taskRepository)

	taskRepository.On("DeleteTaskById", ctx, task.Id).Return(nil)

	err := usecase.DeleteTaskById(ctx, task.Id)

	assert.Equal(t, nil, err)
}
