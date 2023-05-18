package app

import (
	"github.com/ManuelP84/calendar/infra/postgres"
	"github.com/kelseyhightower/envconfig"
)

var instance *AppSettings

type AppSettings struct {
	Port     uint64 `envconfig:"PORT" required:"true"`
	Postgres *postgres.PostgresDbSettings
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
