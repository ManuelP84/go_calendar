package usecase

import (
	"context"
	"fmt"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
)

type DeleteTaskById struct {
	TaskRepository repositories.TaskRepository
}

func NewDeleteTaskById(taskRepository repositories.TaskRepository) *DeleteTaskById {
	return &DeleteTaskById{taskRepository}
}

func (usecase *DeleteTaskById) DeleteTaskById(ctx context.Context, id string) error {
	if id == emptyString {
		return fmt.Errorf("id can't be empty")
	}
	return usecase.TaskRepository.DeleteTaskById(ctx, id)
}
