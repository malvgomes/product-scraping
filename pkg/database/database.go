package database

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	SelectOne(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
}

type dbWrapper struct {
	dbMap *gorp.DbMap
}

func (d *dbWrapper) SelectOne(i interface{}, s string, args ...interface{}) error {
	err := d.dbMap.SelectOne(i, s, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (d *dbWrapper) Exec(s string, args ...interface{}) (sql.Result, error) {
	result, err := d.dbMap.Exec(s, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return result, nil
}

func Open() (Database, error) {
	conn, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/scraper?loc=Local&parseTime=true&charset=utf8mb4")
	if err != nil {
		return nil, err
	}

	dbMap := &gorp.DbMap{
		Db:      conn,
		Dialect: gorp.MySQLDialect{},
	}

	return &dbWrapper{dbMap: dbMap}, nil
}
