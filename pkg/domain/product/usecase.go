package product

import (
	"fmt"
	"log"
	"net/url"
	"product-scraping/pkg/scraper"
	"strconv"
	"time"
)

type usecase struct {
	repository Repository
	scraper    scraper.Scraper
}

type UseCase interface {
	ParseProductURL(url string) (*Entity, error)
}

func NewUseCase(repository Repository, s scraper.Scraper) UseCase {
	return &usecase{
		repository: repository,
		scraper:    s,
	}
}

func (u *usecase) ParseProductURL(productURL string) (*Entity, error) {
	parsedURL, err := url.Parse(productURL)
	if err != nil {
		return nil, err
	}

	productURL = fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, parsedURL.Path)

	data, err := u.repository.GetProductData(productURL)
	if err != nil {
		return nil, err
	}

	if !data.IsEmpty() && data.InsertionDate.After(time.Now().Add(-time.Hour)) {
		log.Println("Selecting data from database")
		return data, nil
	}

	log.Println("Scraping data from url")
	title, image, price, description, err := u.scraper.ScrapeURL(parsedURL.Host, productURL)
	if err != nil {
		return nil, err
	}

	intPrice, err := strconv.Atoi(price)
	if err != nil {
		return nil, err
	}

	ent := &Entity{
		Title:       title,
		ImageURL:    image,
		Price:       intPrice,
		Description: description,
		URL:         productURL,
	}

	log.Println("Inserting data into the database")
	err = u.repository.InsertProductData(ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}
