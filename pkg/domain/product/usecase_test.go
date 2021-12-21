package product_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	productmock "product-scraping/mock/product"
	scrapermock "product-scraping/mock/scraper"
	"product-scraping/pkg/domain/product"
	"testing"
	"time"
)

func TestUsecase_ParseProductURL(t *testing.T) {
	t.Run("Invalid URL", func(t *testing.T) {
		usecase, _, _, finish := getMockedUseCase(t)
		defer finish()

		entity, err := usecase.ParseProductURL("!@#$%*()<><,.")
		assert.Error(t, err)
		assert.Nil(t, entity)
	})

	t.Run("Success - Data from database", func(t *testing.T) {
		usecase, repositoryMock, _, finish := getMockedUseCase(t)
		defer finish()

		ent := &product.Entity{
			Title:         "title",
			ImageURL:      "url",
			Price:         1,
			Description:   "desc",
			URL:           "url",
			InsertionDate: time.Now(),
		}

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(ent, nil)

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.Nil(t, err)
		assert.Equal(t, ent, entity)
	})

	t.Run("Failure - SELECT error", func(t *testing.T) {
		usecase, repositoryMock, _, finish := getMockedUseCase(t)
		defer finish()

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(nil, errors.New("failure"))

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.EqualError(t, err, "failure")
		assert.Nil(t, entity)
	})

	t.Run("Success - Data from scrape", func(t *testing.T) {
		usecase, repositoryMock, scraperMock, finish := getMockedUseCase(t)
		defer finish()

		ent := &product.Entity{
			Title:       "title",
			ImageURL:    "url",
			Price:       1,
			Description: "desc",
			URL:         "https://url.com",
		}

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(nil, nil)
		scraperMock.EXPECT().ScrapeURL("url.com", "https://url.com").
			Return("title", "url", "1", "desc", nil)
		repositoryMock.EXPECT().InsertProductData(&product.Entity{
			Title:       "title",
			ImageURL:    "url",
			Price:       1,
			Description: "desc",
			URL:         "https://url.com",
		}).Return(nil)

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.Nil(t, err)
		assert.Equal(t, ent, entity)
	})

	t.Run("Failure - Scraping data", func(t *testing.T) {
		usecase, repositoryMock, scraperMock, finish := getMockedUseCase(t)
		defer finish()

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(nil, nil)
		scraperMock.EXPECT().ScrapeURL("url.com", "https://url.com").
			Return("", "", "", "", errors.New("failure"))

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.EqualError(t, err, "failure")
		assert.Nil(t, entity)
	})

	t.Run("Failure - Atoi", func(t *testing.T) {
		usecase, repositoryMock, scraperMock, finish := getMockedUseCase(t)
		defer finish()

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(nil, nil)
		scraperMock.EXPECT().ScrapeURL("url.com", "https://url.com").
			Return("title", "url", "batata", "desc", nil)

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.Error(t, err)
		assert.Nil(t, entity)
	})

	t.Run("Failure - INSERT error", func(t *testing.T) {
		usecase, repositoryMock, scraperMock, finish := getMockedUseCase(t)
		defer finish()

		repositoryMock.EXPECT().GetProductData("https://url.com").
			Return(nil, nil)
		scraperMock.EXPECT().ScrapeURL("url.com", "https://url.com").
			Return("title", "url", "1", "desc", nil)
		repositoryMock.EXPECT().InsertProductData(&product.Entity{
			Title:       "title",
			ImageURL:    "url",
			Price:       1,
			Description: "desc",
			URL:         "https://url.com",
		}).Return(errors.New("failure"))

		entity, err := usecase.ParseProductURL("https://url.com")
		assert.EqualError(t, err, "failure")
		assert.Nil(t, entity)
	})
}

func getMockedUseCase(t *testing.T) (product.UseCase, *productmock.MockRepository, *scrapermock.MockScraper, func()) {
	t.Helper()

	ctrl := gomock.NewController(t)

	repositoryMock := productmock.NewMockRepository(ctrl)
	scraperMock := scrapermock.NewMockScraper(ctrl)

	finish := func() {
		ctrl.Finish()
	}

	return product.NewUseCase(repositoryMock, scraperMock), repositoryMock, scraperMock, finish
}
