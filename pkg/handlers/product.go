package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"product-scraping/pkg/domain/product"
)

type ProductHandler struct {
	productUseCase product.UseCase
}

func NewProductHandler(puc product.UseCase) *ProductHandler {
	return &ProductHandler{productUseCase: puc}
}

type payload struct {
	URL string `json:"url"`
}

func (p *payload) IsValid() bool {
	return p.URL != ""
}

// ParseProduct trata o payload, verificando se é válido e retorna um json dos dados já tratados pelo UseCase
func (p *ProductHandler) ParseProduct(w http.ResponseWriter, r *http.Request) {
	var pl payload

	err := json.NewDecoder(r.Body).Decode(&pl)
	if err != nil || !pl.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := p.productUseCase.ParseProductURL(pl.URL)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	content, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
