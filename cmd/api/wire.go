//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ManuelP84/calendar/infra/app"
	"github.com/ManuelP84/calendar/infra/http/server"

	"github.com/google/wire"
)

func CreateApp() *app.App {
	wire.Build(
		app.GetAppSettings,
		server.NewWebServer,
		app.NewApp,
	)
	return new(app.App)
}
