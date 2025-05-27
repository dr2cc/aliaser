package handlers

import (
	"aliaser/internal/config"
	"aliaser/internal/lib/random"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

const aliasLength = 6

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=URLSaver
type URLSaver interface {
	SaveURL(URL, alias string) (int64, error)
}

// 28.05 начать тут!!
// Функция PostHandler уровня пакета handlers
func PostHandler(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			param, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Преобразуем тело запроса (тип []byte) в строку:
			url := string(param)

			// // Генерируем короткий идентификатор и создаем запись в нашем хранилище
			// //config.FlagURL соответствует "http://" + req.Host если не использовать аргументы
			alias := random.NewRandomString(aliasLength)

			// Объект urlSaver (переданный при создании хендлера из main)
			// используется именно тут!
			id, err := urlSaver.SaveURL(url, alias)

			if err != nil {
				fmt.Println("failed to add url")
				return
			}

			// возвращаем ответ с сообщением об успехе
			// Это калька, в таком виде он не нужен
			fmt.Println("url added, id= ", id)

			// Устанавливаем статус ответа 201
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, config.FlagURL+"/"+alias)

		default:
			w.Header().Set("Location", "Method not allowed")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// моя выдумка
type URLGeter interface {
	GetURL(alias string) (string, error)
}

// Функция GetHandler уровня пакета handlers
func GetHandler(urlGeter URLGeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			alias := strings.TrimPrefix(r.RequestURI, "/")

			//url, err := maps.GetEntry(us, id)
			url, err := urlGeter.GetURL(alias)
			if err != nil {
				w.Header().Set("Location", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.Header().Set("Location", url)
			w.WriteHeader(http.StatusTemporaryRedirect)
		default:
			w.Header().Set("Location", "Method not allowed")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
