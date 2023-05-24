package usecase

import (
	"context"
	"testing"

	repoMocks "github.com/ManuelP84/calendar/domain/task/gateways/repositories/mocks"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTask(t *testing.T) {
	ctx := context.TODO()
	taskUpdated := &models.Task{
		Id:          "1",
		Title:       "First task updated",
		Description: "First task description updated",
	}
	taskRepository := repoMocks.NewTaskRepository(t)

	usecase := NewUpdateTask(taskRepository)

	taskRepository.On("UpdateTask", ctx, taskUpdated).Return(nil)

	err := usecase.UpdateTask(ctx, taskUpdated)

	assert.Nil(t, err)
}
