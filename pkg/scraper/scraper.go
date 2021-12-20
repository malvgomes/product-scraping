package scraper

import (
	_ "embed"
	"encoding/json"
	"github.com/gocolly/colly"
	"product-scraping/pkg/domain/product"
)

//go:embed scraper-config.json
var config []byte

type Scraper interface {
	ScrapeURL(url string) (*product.Entity, error)
}

func NewScraper() (Scraper, error) {
	var s scraper
	err := json.Unmarshal(config, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

type scraper map[string]struct {
	ImageTag []string `json:"image_tag"`
}

func (s *scraper) ScrapeURL(url string) (*product.Entity, error) {
	c := colly.NewCollector()

	return nil, nil
}
