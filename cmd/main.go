package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"product-scraping/pkg/database"
	"product-scraping/pkg/domain/product"
	"product-scraping/pkg/handlers"
	"product-scraping/pkg/scraper"
)

// Inicializa todos os recursos utilizados pela api: db, scraper, repository e usecase, e inicia um handler na porta 3000
func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}

	scraperLibrary, err := scraper.NewScraper()
	if err != nil {
		log.Fatal(err)
	}

	productRepository := product.NewRepository(db)

	productUseCase := product.NewUseCase(productRepository, scraperLibrary)

	router := chi.NewRouter()

	// Retorna HTTP 500 em caso de panics, em vez de encerrar a aplicação
	router.Use(middleware.Recoverer)

	// Cria logs da requisição: código de resposta, tempo de execução e tamanho da resposta
	router.Use(middleware.Logger)

	productHandler := handlers.NewProductHandler(productUseCase)

	// Registra o endpoint `product`, e o mapeia para o método ParseProduct do handler
	router.Post("/product", productHandler.ParseProduct)

	log.Println("Listening on port :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
