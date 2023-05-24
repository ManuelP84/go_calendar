package usecase

import (
	"context"
	"fmt"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

type UpdateTask struct {
	TaskRepository repositories.TaskRepository
}

func NewUpdateTask(taskRepository repositories.TaskRepository) *UpdateTask {
	return &UpdateTask{taskRepository}
}

func (usecase *UpdateTask) UpdateTask(ctx context.Context, task *models.Task) error {
	if task == nil {
		return fmt.Errorf("task can't be null")
	}

	if task.Id == "" {
		return fmt.Errorf("task id must be provided")
	}
	return usecase.TaskRepository.UpdateTask(ctx, task)
}
