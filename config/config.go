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
		Host     string `env:"DATABASE_HOST,required"`
		Port     int    `env:"DATABASE_PORT,default=5432"`
		User     string `env:"DATABASE_USER,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		DbName   string `env:"DATABASE_DB_NAME,required"`
	}
	Migrate bool `env:"MIGRATE,default=false"`
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

// PgConn the connection string to the pg database
func PgConn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DbName)
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

// EnableMigrations get enable migrations
func EnableMigrations() bool {
	return c.Migrate
}
