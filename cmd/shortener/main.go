package main

import (
	"aliaser/internal/handlers"
	"aliaser/internal/storage"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	storageInstance := storage.NewStorage()

	mux.HandleFunc("POST /{$}", handlers.PostHandler(storageInstance))
	mux.HandleFunc("GET /{id}", handlers.GetHandler(storageInstance))

	http.ListenAndServe(":8080", mux)
}
