package server

import (
	"flag"
)

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string

// // экспортируемая переменная, отвечает за базовый адрес результирующего сокращённого URL
// var FlagURL string

// parseFlags обрабатывает аргументы командной строки
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
	// регистрируем переменную flagRunAddr
	// как аргумент -a со значением :8080 по умолчанию
	flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	//flag.StringVar(&FlagURL, "b", "http://localhost:8080", "host and port")

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
