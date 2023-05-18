package task

import (
	"github.com/ManuelP84/calendar/infra/app"
)

func RegisterRoutes(app *app.App) *app.App {
	taskGroup := app.WebServer.Group("/task")

	taskGroup.GET("/all", getAllTasks(app.TaskUsecases.GetAllTasks))

	return app
}
