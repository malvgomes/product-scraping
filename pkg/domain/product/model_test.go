package product_test

import (
	"github.com/stretchr/testify/assert"
	"product-scraping/pkg/domain/product"
	"testing"
)

func TestEntity_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		ent   *product.Entity
		empty bool
	}{
		{"Empty", &product.Entity{}, true},
		{"Nil", nil, true},
		{"Not empty", &product.Entity{
			Title:       "Pudim",
			ImageURL:    "http://www.pudim.com.br/pudim.jpg",
			Price:       0,
			Description: "pudim@pudim.com.br ",
			URL:         "http://www.pudim.com.br/",
		}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.empty, test.ent.IsEmpty())
		})
	}
}
