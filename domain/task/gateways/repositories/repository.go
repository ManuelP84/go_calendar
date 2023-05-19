package repositories

import (
	"context"

	"github.com/ManuelP84/calendar/domain/task/models"
)

type TaskRepository interface {
	Close()
	InsertTask(ctx context.Context, task *models.Task) error
	GetTasks(ctx context.Context) ([]*models.Task, error)
	DeleteTaskById(ctx context.Context, id string) error
}
