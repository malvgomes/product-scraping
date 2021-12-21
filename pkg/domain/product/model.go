package product

import "time"

type Entity struct {
	Title         string    `json:"title"`
	ImageURL      string    `json:"imageURL"`
	Price         int       `json:"price"`
	Description   string    `json:"description"`
	URL           string    `json:"url"`
	InsertionDate time.Time `json:"-"`
}

func (e *Entity) IsEmpty() bool {
	return e == nil || *e == Entity{}
}
