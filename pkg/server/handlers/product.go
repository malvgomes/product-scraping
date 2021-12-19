package handlers

import (
	"encoding/json"
	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return new(ProductHandler)
}

type payload struct {
	URL string
}

func (p *payload) IsValid() bool {
	return p.URL != ""
}

func (p *ProductHandler) ParseProduct(w http.ResponseWriter, r *http.Request) {
	var pl payload

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil || !pl.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
