package config

type Config struct {
  ServerPort        string    `env:"SERVER_PORT`
}

type Redis struct {
	ConnectionURL       string   `env:"REDIS_CONNECTION_URL"`
	Password            string   `env:"REDIS_PASSWORD" json:"-"`
	DB                  int      `env:"REDIS_DB"`
	PoolSize            int      `env:"REDIS_POOL_SIZE"`
  PoolTimeout         int      `env:"REDIS_POOL_TIMEOUT"`
}

type Postgres struct {
	PostgresUser     string       `env:"POSTGRES_USER"`
	PostgresPassword string       `env:"POSTGRES_PASSWORD"`
	PostgresHost     string       `env:"POSTGRES_HOST"`
	PostgresPort     string       `env:"POSTGRES_PORT"`
	PostgresDatabase string       `env:"POSTGRES_DATABASE"`
}

var config Config