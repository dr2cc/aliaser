package main

import "github.com/dr2cc/URLsShortener.git/internal/server"

func main() {
	// //mux := http.NewServeMux()
	// mux := chi.NewRouter()

	// storageInstance := storage.NewStorage()

	// // // Обработчик net/http
	// //mux.HandleFunc("POST /{$}", handlers.PostHandler(storageInstance))
	// //mux.HandleFunc("GET /{id}", handlers.GetHandler(storageInstance))

	// // Обработчик chi
	// mux.Post("/", handlers.PostHandler(storageInstance))
	// mux.Get("/{id}", handlers.GetHandler(storageInstance))

	// http.ListenAndServe(":8080", mux)

	// обрабатываем аргументы командной строки
	server.ParseFlags()

	if err := server.Run(); err != nil {
		panic(err)
	}
}
