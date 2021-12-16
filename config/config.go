package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Server struct {
		HTTPPort int32 `env:"HTTPPORT,default=3000"`
	}

	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT,default=5432"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		DbName   string `env:"DATABASE_DB_NAME"`
	}
}

var c config

const (
	//IMGPrefix is the imagen prefix
	IMGPrefix string = "IMG"
)

//ReadConfig read config
func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

// HTTPListener the listener string for the http service
func HTTPListener() string {
	return fmt.Sprintf(":%d", c.Server.HTTPPort)
}

// PgConnMigration returns the config string for migration
func PgConnMigration() *string {
	if c.Migrate {
		pgconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DbName)
		return &pgconn
	}

	return nil
}
