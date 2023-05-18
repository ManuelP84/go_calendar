//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ManuelP84/calendar/business/task/usecase"
	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/infra/app"
	"github.com/ManuelP84/calendar/infra/http/server"
	"github.com/ManuelP84/calendar/infra/postgres"

	"github.com/google/wire"
)

func CreateApp() *app.App {
	wire.Build(
		app.GetAppSettings,
		app.GetPostgresBdSettings,
		server.NewWebServer,
		postgres.NewPostgresRepository,
		usecase.NewTaskUsecases,
		wire.Bind(new(repositories.TaskRepository), new(*postgres.PostgresRepository)),
		app.NewApp,
	)
	return new(app.App)
}
