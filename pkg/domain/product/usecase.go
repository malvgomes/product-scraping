package product

import (
	"net/url"
	"time"
)

type usecase struct {
	repository Repository
}

type UseCase interface {
	ParseProductURL(url string) (*Entity, error)
}

func NewUseCase(repository Repository) UseCase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) ParseProductURL(productURL string) (*Entity, error) {
	data, err := u.repository.GetProductData(productURL)
	if err != nil {
		return nil, err
	}

	if !data.IsEmpty() && data.InsertionDate.After(time.Now().Add(-time.Hour)) {
		return data, nil
	}

	parsedURL, err := url.Parse(productURL)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
