package app

var instance *AppSettings

type AppSettings struct {
	Port uint64 `envconfig:"PORT" required:"true"`
}

func loadAppSettings() *AppSettings {
	if instance == nil {
		settings := AppSettings{
			Port: 8080,
		}
		// if err := envconfig.Process("", &settings); err != nil {
		// 	panic(err)
		// }

		instance = &settings
	}

	return instance
}

func GetAppSettings() *AppSettings {
	return loadAppSettings()
}
