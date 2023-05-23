package usecase

import (
	"context"
	"fmt"

	"github.com/ManuelP84/calendar/domain/task/gateways/bus"
	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/domain/task/models"
)

const (
	emptyString = ""
)

type InsertTask struct {
	TaskRepository repositories.TaskRepository
	TaskBus        bus.TaskBus
}

func NewInsertTask(taskRepository repositories.TaskRepository, taskBus bus.TaskBus) *InsertTask {
	return &InsertTask{taskRepository, taskBus}
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
	usecase.TaskRepository.InsertTask(ctx, task)

	return usecase.TaskBus.Publish("taskCreated")
}
