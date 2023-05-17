package usecase

import (
	"context"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

type GetTasks struct {
	TaskRepository repositories.TaskRepository
}

func NewGetTasks(taskRepository repositories.TaskRepository) *GetTasks {
	return &GetTasks{taskRepository}
}

func (usecase *GetTasks) GetTasks(ctx context.Context) ([]*models.Task, error) {
	return usecase.TaskRepository.GetTasks(ctx)
}
