package handlers

import (
	"github.com/ManuelP84/calendar/infra/app"
	"github.com/ManuelP84/calendar/infra/http/handlers/task"
)

func SetRoutes(app *app.App) {
	task.RegisterRoutes(app)
}
