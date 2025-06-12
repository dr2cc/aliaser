package main

import (
	"aliaser/internal/config"
	"aliaser/internal/http-server/handlers"
	"aliaser/internal/storage/maps"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
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
	router := chi.NewRouter()

	// Примитивное хранилище - map
	storageInstance := maps.NewStorage()

	// // sqlite.New или "подключает" файл db , а если его нет то создает
	// storageInstance, err := sqlite.New("./storage.db")
	// if err != nil {
	// 	//log.Error("failed to initialize storage", sl.Err(err))
	// 	fmt.Println("failed to initialize storage")
	// 	//errors.New("failed to initialize storage")
	// }
	// //

	router.Post("/", handlers.PostHandler(storageInstance))
	router.Get("/{id}", handlers.GetHandler(storageInstance))

	//fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(config.FlagRunAddr, router)
}
