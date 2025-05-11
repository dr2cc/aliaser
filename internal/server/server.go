package server

import (
	"net/http"

	"github.com/dr2cc/URLsShortener.git/internal/handlers"
	"github.com/dr2cc/URLsShortener.git/internal/storage"
	"github.com/go-chi/chi"
)

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func Run() error {
	mux := chi.NewRouter()

	storageInstance := storage.NewStorage()

	mux.Post("/", handlers.PostHandler(storageInstance))
	mux.Get("/{id}", handlers.GetHandler(storageInstance))

	//fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(flagRunAddr, mux)
}
