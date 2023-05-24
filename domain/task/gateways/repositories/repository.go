package repositories

import (
	"context"

	"github.com/ManuelP84/calendar/domain/task/models"
)

//go:generate mockery --name TaskRepository
type TaskRepository interface {
	Close()
	InsertTask(ctx context.Context, task *models.Task) error
	GetTasks(ctx context.Context) ([]*models.Task, error)
	DeleteTaskById(ctx context.Context, id string) error
	SearchTaskById(ctx context.Context, id string) (*models.Task, error)
}
