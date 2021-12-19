package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"product-scraping/pkg/server/handlers"
)

type Router struct{}

func NewRouter() *Router {

	return new(Router)
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	productHandler := handlers.NewProductHandler()

	router.Post("/product", productHandler.ParseProduct)

	return router
}
