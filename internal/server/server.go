package server

import (
	"aliaser/internal/config"
	"aliaser/internal/handlers"
	"aliaser/internal/storage"
	"net/http"

	"github.com/go-chi/chi"
)

// инициализации зависимостей сервера перед запуском
func Run() error {
	mux := chi.NewRouter()

	storageInstance := storage.NewStorage()

	mux.Post("/", handlers.PostHandler(storageInstance))
	mux.Get("/{id}", handlers.GetHandler(storageInstance))

	//fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(config.FlagRunAddr, mux)
}
