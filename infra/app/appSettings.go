package app

import (
	"github.com/ManuelP84/calendar/infra/postgres"
	"github.com/ManuelP84/calendar/infra/rabbit"
	"github.com/kelseyhightower/envconfig"
)

var instance *AppSettings

type AppSettings struct {
	Port     uint64 `envconfig:"PORT" required:"true"`
	Postgres *postgres.PostgresDbSettings
	Rabbit   *rabbit.RabbitSettings
}

func loadAppSettings() *AppSettings {
	if instance == nil {
		settings := AppSettings{}

		if err := envconfig.Process("", &settings); err != nil {
			panic(err)
		}

		instance = &settings
	}

	return instance
}

func GetAppSettings() *AppSettings {
	return loadAppSettings()
}

func GetPostgresBdSettings() *postgres.PostgresDbSettings {
	return loadAppSettings().Postgres
}

func GetRabbitSettings() *rabbit.RabbitSettings {
	return loadAppSettings().Rabbit
}
