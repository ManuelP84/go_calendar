package app

import (
	"context"
	"fmt"
	"log"

	"github.com/ManuelP84/calendar/business/task/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type App struct {
	WebServer    *gin.Engine
	Settings     *AppSettings
	TaskUsecases *usecase.TaskUsecases
}

func NewApp(webServer *gin.Engine, settings *AppSettings, taskUsecases *usecase.TaskUsecases) *App {
	return &App{
		WebServer:    webServer,
		Settings:     settings,
		TaskUsecases: taskUsecases,
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

	log.Printf("WebServer running on port: %s\n", port)

	return g.Wait()
}
