package main

import (
	"net/http"

	"github.com/dr2cc/URLsShortener.git/internal/config"
	"github.com/dr2cc/URLsShortener.git/internal/http-server/handlers"
	maps "github.com/dr2cc/URLsShortener.git/internal/storage/maps"
	"github.com/go-chi/chi"
)

func main() {
	// mux := http.NewServeMux()
	// storageInstance := storage.NewStorage()

	// mux.HandleFunc("POST /{$}", handlers.PostHandler(storageInstance))
	// mux.HandleFunc("GET /{id}", handlers.GetHandler(storageInstance))

	// http.ListenAndServe(":8080", mux)

	// обрабатываем аргументы командной строки
	config.ParseFlags()

	if err := Run(); err != nil {
		panic(err)
	}

	// Объявить переменные окружения так:
	// $env:SERVER_ADDRESS = "localhost:8089"
	// $env:BASE_URL  = "http://localhost:9999"
}

// инициализации зависимостей сервера перед запуском
func Run() error {
	mux := chi.NewRouter()
	storageInstance := maps.NewStorage()

	mux.Post("/", handlers.PostHandler(storageInstance))
	mux.Get("/{id}", handlers.GetHandler(storageInstance))

	//fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(config.FlagRunAddr, mux)
}
