package database_mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-gorp/gorp"
	"product-scraping/pkg/database"
)

func NewDBMock() (database.Database, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return &database.DbWrapper{DbMap: &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}}, mock, nil
}
