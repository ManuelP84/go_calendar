package usecase

import (
	"context"
	"fmt"

	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

const (
	emptyString = ""
)

type InsertTask struct {
	TaskRepository repositories.TaskRepository
}

func NewInsertTask(taskRepository repositories.TaskRepository) *InsertTask {
	return &InsertTask{taskRepository}
}

func (usecase *InsertTask) InsertTask(ctx context.Context, task *models.Task) error {
	// Some validations
	if task == nil {
		return fmt.Errorf("task can't be null")
	}
	if task.Title == emptyString {
		return fmt.Errorf("title can't be empty")
	}
	if task.Description == emptyString {
		return fmt.Errorf("description can't be empty")
	}
	return usecase.TaskRepository.InsertTask(ctx, task)
}
