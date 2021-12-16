package datastore

import (
	"database/sql"

	"github.com/ansel1/merry"
	"github.com/doug-martin/goqu/v9"
)

//SQLDB connection database
var (
	SQLDB     *sql.DB
	SQLDBGoqu goqu.Database
)

//NewDBConn Connects to postgress database
func NewDBConn() error {
	psqlInfo := config.PgConn()
	return newDB(psqlInfo)
}

func newDB(psqlInfo string) error {
	var err error
	SQLDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {

		return merry.Wrap(err)
	}
	err = SQLDB.Ping()
	if err != nil {

		return merry.Wrap(err)
	}

	SQLDBGoqu = *goqu.Dialect("postgres").DB(SQLDB)

	return nil
}
