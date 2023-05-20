package app

import (
	"context"
	"fmt"
	"log"

	"github.com/ManuelP84/calendar/business/task/usecase"
	"github.com/ManuelP84/calendar/infra/rabbit/producer/task"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type App struct {
	WebServer *gin.Engine
	Settings  *AppSettings

	TaskUsecases *usecase.TaskUsecases
	TaskBus      *task.TaskProducer
}

func NewApp(webServer *gin.Engine, settings *AppSettings, taskUsecases *usecase.TaskUsecases, taskBus *task.TaskProducer) *App {
	return &App{
		WebServer:    webServer,
		Settings:     settings,
		TaskUsecases: taskUsecases,
		TaskBus:      taskBus,
	}
}

func (app *App) Run(ctx context.Context) error {
	port := fmt.Sprintf(":%d", app.Settings.Port)
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := app.WebServer.Run(port); err != nil {
			return err
		}
		return nil
	})

	log.Printf("WebServer running on port %s\n", port)

	return g.Wait()
}
