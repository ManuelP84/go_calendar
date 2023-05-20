//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ManuelP84/calendar/business/task/usecase"
	//"github.com/ManuelP84/calendar/domain/task/gateways/bus"
	"github.com/ManuelP84/calendar/domain/task/gateways/repositories"
	"github.com/ManuelP84/calendar/infra/app"
	"github.com/ManuelP84/calendar/infra/http/server"
	"github.com/ManuelP84/calendar/infra/postgres"
	"github.com/ManuelP84/calendar/infra/rabbit/producer/task"

	"github.com/google/wire"
)

func CreateApp() *app.App {
	wire.Build(
		app.GetAppSettings,
		app.GetPostgresBdSettings,
		app.GetRabbitSettings,
		server.NewWebServer,
		postgres.NewPostgresRepository,
		task.NewTaskProducer,
		usecase.NewTaskUsecases,
		wire.Bind(new(repositories.TaskRepository), new(*postgres.PostgresRepository)),
		//wire.Bind(new(bus.TaskBus), new(*task.TaskProducer)),
		app.NewApp,
	)
	return new(app.App)
}
