package usecase

import (
	"context"
	"testing"

	repoMocks "github.com/ManuelP84/calendar/domain/task/gateways/repositories/mocks"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/stretchr/testify/assert"
)

func TestSearchTaskById(t *testing.T) {
	ctx := context.TODO()
	task := &models.Task{
		Id:          "1",
		Title:       "First task",
		Description: "First task description",
	}

	taskRepository := repoMocks.NewTaskRepository(t)

	usecase := NewSearchTaskById(taskRepository)

	taskRepository.On("SearchTaskById", ctx, task.Id).Return(task, nil)

	searchedTask, err := usecase.SearchTaskById(ctx, task.Id)

	assert.Equal(t, searchedTask, task)
	assert.Nil(t, err)
}
