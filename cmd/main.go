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
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	productHandler := handlers.NewProductHandler(productUseCase)

	router.Post("/product", productHandler.ParseProduct)

	log.Println("Listening on port :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
