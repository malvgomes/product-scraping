package product

import (
	_ "embed"
	"github.com/nleof/goyesql"
	"product-scraping/pkg/database"
)

//go:embed queries.sql
var queries []byte

type repository struct {
	db      database.Database
	queries goyesql.Queries
}

type Repository interface {
	InsertProductData(data Entity) error
	GetProductData(url string) (*Entity, error)
}

func NewRepository(db database.Database) Repository {
	return &repository{
		db:      db,
		queries: goyesql.MustParseBytes(queries),
	}
}

func (r *repository) InsertProductData(data Entity) error {
	_, err := r.db.Exec(r.queries["get-data"], data.Title, data.ImageURL, data.Price, data.Description, data.URL)
	return err
}

func (r *repository) GetProductData(url string) (*Entity, error) {
	var data Entity
	err := r.db.SelectOne(&data, r.queries["get-data"], url)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
