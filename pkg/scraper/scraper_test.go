package scraper_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"product-scraping/pkg/scraper"
	"testing"
)

func TestScraper_ScrapeURL(t *testing.T) {
	t.Run("Invalid host", func(t *testing.T) {
		s := getMockedScraper(t)

		title, image, price, description, err := s.ScrapeURL("www.macarrao.com", "https://www.macarrao.com")
		assert.Error(t, err)
		assert.Empty(t, title)
		assert.Empty(t, image)
		assert.Empty(t, price)
		assert.Empty(t, description)
	})

	t.Run("Success", func(t *testing.T) {
		s := getMockedScraper(t)

		httpmock.Activate()
		defer httpmock.Deactivate()

		httpmock.RegisterResponder(http.MethodGet, "https://www.test.com.br",
			func(request *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(http.StatusOK, `
					<!DOCTYPE html>
					<html>
						<head>
							<meta property="title" content="title">
							<div class="batata"><strong class="macarrao">//image.com?q=b</strong></div>
							<div id="feijao">49,99</div>							
							<meta property="description" content="description">
						</head>
						<body>
							<h1>Test page</h1>
						</body>
					</html>
				`)

				resp.Header.Add("Content-Type", "text/html")
				return resp, nil
			},
		)

		title, image, price, description, err := s.ScrapeURL("www.test.com.br", "https://www.test.com.br")
		assert.Nil(t, err)
		assert.Equal(t, "title", title)
		assert.Equal(t, "https://image.com?q=b", image)
		assert.Equal(t, "4999", price)
		assert.Equal(t, "description", description)
	})

	t.Run("Failure - Forbidden", func(t *testing.T) {
		s := getMockedScraper(t)

		httpmock.Activate()
		defer httpmock.Deactivate()

		httpmock.RegisterResponder(http.MethodGet, "https://www.test.com.br",
			httpmock.NewStringResponder(http.StatusForbidden, "Forbidden"),
		)

		title, image, price, description, err := s.ScrapeURL("www.test.com.br", "https://www.test.com.br")
		assert.Error(t, err)
		assert.Empty(t, title)
		assert.Empty(t, image)
		assert.Empty(t, price)
		assert.Empty(t, description)
	})

	t.Run("Failure - Invalid image URL", func(t *testing.T) {
		s := getMockedScraper(t)

		httpmock.Activate()
		defer httpmock.Deactivate()

		httpmock.RegisterResponder(http.MethodGet, "https://www.test.com.br",
			func(request *http.Request) (*http.Response, error) {
				resp := httpmock.NewStringResponse(http.StatusOK, `
					<!DOCTYPE html>
					<html>
						<head>
							<meta property="title" content="title">
							<div class="batata"><strong class="macarrao">!@#$%*()_<>:<</strong></div>
							<div id="feijao">49,99</div>							
							<meta property="description" content="description">
						</head>
						<body>
							<h1>Test page</h1>
						</body>
					</html>
				`)

				resp.Header.Add("Content-Type", "text/html")
				return resp, nil
			},
		)

		title, image, price, description, err := s.ScrapeURL("www.test.com.br", "https://www.test.com.br")
		assert.Error(t, err)
		assert.Empty(t, title)
		assert.Empty(t, image)
		assert.Empty(t, price)
		assert.Empty(t, description)
	})
}

func getMockedScraper(t *testing.T) scraper.Scraper {
	t.Helper()

	s, err := scraper.NewScraper()
	assert.Nil(t, err)

	return s
}
