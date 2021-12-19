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
