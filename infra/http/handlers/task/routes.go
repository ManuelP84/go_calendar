package task

import (
	"github.com/ManuelP84/calendar/app"
)

func RegisterRoutes(app *app.App) *app.App {
	taskGroup := app.WebServer.Group("/task")

	taskGroup.GET("/all", getAllTasks(app.TaskUsecases.GetAllTasks))

	taskGroup.GET("/search/:id", getTaskById(app.TaskUsecases.SearchTaskById))

	taskGroup.POST("/create", insertTask(app.TaskUsecases.InsertTask))

	taskGroup.PUT("/update", updateTask(app.TaskUsecases.UpdateTask))

	taskGroup.DELETE("/delete/:id", deleteTaskById(app.TaskUsecases.DeleteTaskById))
	return app
}
