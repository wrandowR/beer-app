package datastore

import (
	"ZachIgarz/test-beer/config"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/golang-migrate/migrate/v4"

	// blank
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// DoMigration does the migration in postgres
func DoMigration() {
	pgconn := config.PgConnMigration()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = formatDir(dir)
	sourcedir := "file://" + path.Join(dir, "/infrastructure/datastore/migrations")

	migration, err := migrate.New(sourcedir, *pgconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading migration sourcedir=%s pgconn=%s error=%s %v\n", sourcedir, *pgconn, err.Error(), err)
		panic(err)
	}
	defer migration.Close()

	err = migration.Up()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en migracion. %s. data %v\n", err.Error(), err)
	}
}

//ResetDatabase clean database and do migration, execute with precaution, use only in testing or development
func ResetDatabase() {
	undoMigration()
	DoMigration()
}

// undoMigration does the migration in postgres
func undoMigration() {
	pgconn := config.PgConnMigration()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = formatDir(dir)
	sourcedir := "file://" + path.Join(dir, "/infrastructure/datastore/migrations")

	migration, err := migrate.New(sourcedir, *pgconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading migration sourcedir=%s pgconn=%s error=%s %v\n", sourcedir, *pgconn, err.Error(), err)
		panic(err)
	}
	defer migration.Close()

	err = migration.Drop()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error en migracion. %s. data %v\n", err.Error(), err)
	}
}

func formatDir(dir string) string {
	basedir := strings.Split(dir, "/beer-app")[0]
	return basedir + "/beer-app"
}
