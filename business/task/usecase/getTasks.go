package usecase

import (
	"context"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

type GetAllTasks struct {
	TaskRepository repositories.TaskRepository
}

func NewGetAllTasks(taskRepository repositories.TaskRepository) *GetAllTasks {
	return &GetAllTasks{taskRepository}
}

func (usecase *GetAllTasks) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	return usecase.TaskRepository.GetTasks(ctx)
}
