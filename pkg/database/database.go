package database

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Database interface {
	SelectOne(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
}

type DbWrapper struct {
	DbMap *gorp.DbMap
}

func (d *DbWrapper) SelectOne(i interface{}, s string, args ...interface{}) error {
	err := d.DbMap.SelectOne(i, s, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (d *DbWrapper) Exec(s string, args ...interface{}) (sql.Result, error) {
	result, err := d.DbMap.Exec(s, args...)
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

	err = conn.Ping()
	for err != nil {
		log.Println("Database is not yet ready. Trying again")
		time.Sleep(time.Second * 5)
		err = conn.Ping()
	}

	dbMap := &gorp.DbMap{
		Db:      conn,
		Dialect: gorp.MySQLDialect{},
	}

	return &DbWrapper{DbMap: dbMap}, nil
}
