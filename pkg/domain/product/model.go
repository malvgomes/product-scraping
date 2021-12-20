package product

import "time"

type Entity struct {
	Title         string
	ImageURL      string
	Price         int
	Description   string
	URL           string
	InsertionDate time.Time
}

func (e *Entity) IsEmpty() bool {
	return e == nil || *e == Entity{}
}

type ScraperParameters struct {
	Title       string
	ImageURL    string
	Price       int
	Description string
}
