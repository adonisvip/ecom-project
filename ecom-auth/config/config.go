package config

type Config struct {
  // ServerPort        string    `env:"SERVER_PORT`
  Redis             *Redis        
  Postgres          *Postgres      
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


// func LoadConfig(cfg interface{}) error {
// 	v := reflect.ValueOf(cfg).Elem()

// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Field(i)

// 		// Chỉ xử lý nếu là struct
// 		if field.Kind() == reflect.Struct {
// 			fieldAddr := field.Addr().Interface()
// 			if err := envconfig.Process("", fieldAddr); err != nil {
// 				return fmt.Errorf("failed to process field %s: %w", v.Type().Field(i).Name, err)
// 			}
// 		}
// 	}
// 	return nil
// }