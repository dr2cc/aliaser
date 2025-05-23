package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/dr2cc/URLsShortener.git/internal/config"
	maps "github.com/dr2cc/URLsShortener.git/internal/storage/maps"
)

const aliasLength = 6

func generateAlias(us *maps.URLStorage, url string) string {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, aliasLength)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	id := string(b)
	maps.MakeEntry(us, id, url)

	return "/" + id
}

// Функция PostHandler уровня пакета handlers
func PostHandler(us *maps.URLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			param, err := io.ReadAll(req.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Преобразуем тело запроса (тип []byte) в строку:
			url := string(param)
			// // Генерируем короткий идентификатор и создаем запись в нашем хранилище
			//alias := "http://" + req.Host + generateAlias(us, url)

			alias := config.FlagURL + generateAlias(us, url)

			// Устанавливаем статус ответа 201
			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, alias)

		default:
			w.Header().Set("Location", "Method not allowed")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// Функция GetHandler уровня пакета handlers
func GetHandler(us *maps.URLStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			id := strings.TrimPrefix(req.RequestURI, "/")
			url, err := maps.GetEntry(us, id)
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
