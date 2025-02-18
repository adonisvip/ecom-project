package postgres

import (
	"database/sql"
	"fmt"
	cfg "ecom-auth/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresInitConnection(cfg *cfg.Postgres) (*gorm.DB, *sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConfig, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	// dbConfig.SetMaxOpenConns(maxOpenConns)
	// dbConfig.SetConnMaxLifetime(connMaxLifetime * time.Second)
	// dbConfig.SetMaxIdleConns(maxIdleConns)
	// dbConfig.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = dbConfig.Ping(); err != nil {
		return nil, nil, err
	}
	return db, dbConfig, nil
}