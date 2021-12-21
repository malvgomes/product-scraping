package product_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	databasemock "product-scraping/mock/database"
	"product-scraping/pkg/domain/product"
	"regexp"
	"testing"
	"time"
)

func TestRepository_InsertProductData(t *testing.T) {
	query := regexp.QuoteMeta(`
		INSERT INTO scraper.products (title, image_url, price, description, url) VALUES (?,?,?,?,?)
		ON DUPLICATE KEY UPDATE title = VALUES(title), image_url = VALUES(image_url), price = VALUES(price), description = VALUES(description), date = NOW();
	`)

	t.Run("Success", func(t *testing.T) {
		repo, mock := getMockedRepository(t)

		mock.ExpectExec(query).
			WithArgs("Title", "Image", 1, "Description", "URL").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.InsertProductData(&product.Entity{
			Title:       "Title",
			ImageURL:    "Image",
			Price:       1,
			Description: "Description",
			URL:         "URL",
		})
		assert.Nil(t, err)
	})

	t.Run("Failure", func(t *testing.T) {
		repo, mock := getMockedRepository(t)

		mock.ExpectExec(query).
			WithArgs("Title", "Image", 1, "Description", "URL").
			WillReturnError(errors.New("failure"))

		err := repo.InsertProductData(&product.Entity{
			Title:       "Title",
			ImageURL:    "Image",
			Price:       1,
			Description: "Description",
			URL:         "URL",
		})
		assert.EqualError(t, err, "failure")
	})
}

func TestRepository_GetProductData(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT
		    title Title,
		    image_url ImageURL,
		    price Price,
		    description Description,
		    url URL,
		    date InsertionDate
		FROM scraper.products
		WHERE url = ?;
	`)

	t.Run("Success", func(t *testing.T) {
		repo, mock := getMockedRepository(t)

		loc, err := time.LoadLocation("America/Sao_Paulo")
		assert.Nil(t, err)

		date := time.Date(2022, 1, 1, 11, 12, 13, 0, loc)

		mock.ExpectQuery(query).
			WithArgs("URL").
			WillReturnRows(sqlmock.NewRows([]string{"Title", "ImageURL", "Price", "Description", "URL", "InsertionDate"}).
				AddRow("title", "image", 1, "description", "url", date))

		data, err := repo.GetProductData("URL")
		assert.Nil(t, err)
		assert.Equal(t, &product.Entity{
			Title:         "title",
			ImageURL:      "image",
			Price:         1,
			Description:   "description",
			URL:           "url",
			InsertionDate: date,
		}, data)
	})

	t.Run("Failure", func(t *testing.T) {
		repo, mock := getMockedRepository(t)

		mock.ExpectQuery(query).
			WithArgs("URL").
			WillReturnError(errors.New("failure"))

		data, err := repo.GetProductData("URL")
		assert.EqualError(t, err, "failure")
		assert.Nil(t, data)
	})
}

func getMockedRepository(t *testing.T) (product.Repository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := databasemock.NewDBMock()
	assert.Nil(t, err)

	return product.NewRepository(db), mock
}
