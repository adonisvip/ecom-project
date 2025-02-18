package config

type Config struct {
	LogLevel         string `env:"LOG_LEVEL"`
	Environment      string `env:"ENVIRONMENT"`
	ServiceName      string `env:"SERVICE_NAME"`
	ServiceVersion   string `env:"SERVICE_VERSION"`
	BaseURL          string `env:"BASE_URL"`
	HTTPPort         int    `env:"HTTP_PORT"`

	Redis    *Redis
}

type Redis struct {
	ConnectionURL       string   `env:"REDIS_CONNECTION_URL"`
	Password            string   `env:"REDIS_PASSWORD" json:"-"`
	DB                  int      `env:"REDIS_DB"`
	PoolSize            int      `env:"REDIS_POOL_SIZE"`
  PoolTimeout         int      `env:"REDIS_POOL_TIMEOUT"`
	// DialTimeoutSeconds  int      `env:"REDIS_DIAL_TIMEOUT_SECONDS"`
	// ReadTimeoutSeconds  int      `env:"REDIS_READ_TIMEOUT_SECONDS"`
	// WriteTimeoutSeconds int      `env:"REDIS_WRITE_TIMEOUT_SECONDS"`
	// IdleTimeoutSeconds  int      `env:"REDIS_IDLE_TIMEOUT_SECONDS"`
	// MaxIdleConn         int      `env:"REDIS_MAX_IDLE_CONN_NUMBER"`
	// MaxActiveConn       int      `env:"REDIS_MAX_ACTIVE_CONN_NUMBER"`
}

type ServiceGrpc struct {
  AuthGRPC            string    `env:"AUTH_SERVICE_GRPC"`
  CoreGRPC            string    `env:"CORE_SERVICE_GRPC"`
}

var GrpcConfig ServiceGrpc