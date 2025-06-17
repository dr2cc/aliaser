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

	// storages

	// Примитивное (based on map) хранилище
	storageInstance := maps.NewURLStorage(make(map[string]string))

	// // Sqlite хранилище
	// // sqlite.New или "подключает" файл db , а если его нет то создает
	// storageInstance, err := sqlite.New("./storage.db")
	// if err != nil {
	// 	// // Когда подключу логирование- пригодится!
	// 	//log.Error("failed to initialize storage", sl.Err(err))
	// 	fmt.Println("failed to initialize storage")
	// 	//errors.New("failed to initialize storage")
	// }
	// //

	// routers
	//
	// В Go передача интерфейса параметром в функцию означает,
	// что функция может принимать на вход объект любого типа,
	// который реализует определенный интерфейс.
	//
	// PostHandler принимает параметром интерфейс URLSaver
	// с единственным методом SaveURL(URL, alias string) error
	// т.е. два строковых значения .
	// НО! Самое важное- то, что мы передадим параметром должно
	// реализовывать МЕТОДЫ интерфейса!
	router.Post("/", handlers.PostHandler(storageInstance))
	// GetHandler принимает ...
	router.Get("/{id}", handlers.GetHandler(storageInstance))

	// server
	//fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(config.FlagRunAddr, router)
}
