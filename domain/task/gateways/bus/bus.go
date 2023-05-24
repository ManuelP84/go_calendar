package bus

import "github.com/ManuelP84/calendar/domain/task/events"

//go:generate mockery --name TaskBus
type TaskBus interface {
	Publish(event events.TaskEvent) error
}
