package usecase

import (
	"github.com/ManuelP84/calendar/domain/task/gateways/bus"
	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
)

type TaskUsecases struct {
	GetAllTasks    *GetAllTasks
	InsertTask     *InsertTask
	DeleteTaskById *DeleteTaskById
	SearchTaskById *SearchTaskById
	UpdateTask     *UpdateTask
}

func NewTaskUsecases(taskRepository repositories.TaskRepository, taskBus bus.TaskBus) *TaskUsecases {
	return &TaskUsecases{
		GetAllTasks:    NewGetAllTasks(taskRepository),
		InsertTask:     NewInsertTask(taskRepository, taskBus),
		DeleteTaskById: NewDeleteTaskById(taskRepository),
		SearchTaskById: NewSearchTaskById(taskRepository),
		UpdateTask:     NewUpdateTask(taskRepository),
	}
}
