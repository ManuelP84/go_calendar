package usecase

import (
	"context"
	"fmt"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

type SearchTaskById struct {
	TaskRepository repositories.TaskRepository
}

func NewSearchTaskById(taskRepository repositories.TaskRepository) *SearchTaskById {
	return &SearchTaskById{taskRepository}
}

func (usecase *SearchTaskById) SearchTaskById(ctx context.Context, id string) (*models.Task, error) {
	if id == emptyString {
		return nil, fmt.Errorf("id can't be empty")
	}
	return usecase.TaskRepository.SearchTaskById(ctx, id)
}
