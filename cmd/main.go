package main

import (
	"log"
	"net/http"
	"product-scraping/pkg/server"
)

func main() {
	router := server.NewRouter().Init()

	log.Println("Listening on port :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
