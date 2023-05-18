package postgres

type PostgresDbSettings struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Db       string `envconfig:"POSTGRES_DB" required:"true"`
}
