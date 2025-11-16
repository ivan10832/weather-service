package main

import (
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("weather sss"))
		if err != nil {
			log.Printf("error writing response: %v", err)
		}
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Printf("error starting server: %v", err)

	}
}
