package usecase

import (
	"context"
	"testing"

	repoMocks "github.com/ManuelP84/calendar/domain/task/gateways/repositories/mocks"
	"github.com/ManuelP84/calendar/domain/task/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	ctx := context.TODO()

	var repoTasks []*models.Task

	t1 := &models.Task{

		Id:          "1",
		Title:       "First task",
		Description: "First task description",
	}

	t2 := &models.Task{
		Id:          "2",
		Title:       "Second task",
		Description: "Second task description",
	}

	repoTasks = append(repoTasks, t1, t2)

	taskRepository := repoMocks.NewTaskRepository(t)

	usecase := NewGetAllTasks(taskRepository)

	taskRepository.On("GetTasks", ctx).Return(repoTasks, nil)

	tasks, err := usecase.GetAllTasks(ctx)

	assert.Equal(t, tasks, repoTasks)
	assert.ErrorIs(t, err, nil)
}
