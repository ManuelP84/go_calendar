package handlers

import (
	"github.com/ManuelP84/calendar/app"
	"github.com/ManuelP84/calendar/infra/http/handlers/task"
)

func SetRoutes(app *app.App) {
	task.RegisterRoutes(app)
}
