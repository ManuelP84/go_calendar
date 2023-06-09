package events

import "github.com/ManuelP84/calendar/domain/task/models"

const (
	TaskCreatedEvent = "taskCreated"
)

type TaskEvent struct {
	EventType string
	Task      *models.Task
}
