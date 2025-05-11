package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dr2cc/URLsShortener.git/internal/storage"
)

func TestGetHandler(t *testing.T) {
	//Здесь общие для всех тестов данные
	shortURL := "6ba7b811"
	record := map[string]string{shortURL: "https://practicum.yandex.ru/"}

	tests := []struct {
		name       string
		method     string
		input      *storage.URLStorage
		want       string
		wantStatus int
	}{
		{
			name:   "all good",
			method: http.MethodGet,
			input: &storage.URLStorage{
				Data: record,
			},
			want:       "https://practicum.yandex.ru/",
			wantStatus: http.StatusTemporaryRedirect,
		},
		{
			name:   "with bad method",
			method: http.MethodPost,
			input: &storage.URLStorage{
				Data: record,
			},
			want:       "Method not allowed",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:   "key in input does not match /6ba7b811",
			method: http.MethodGet,
			input: &storage.URLStorage{
				Data: map[string]string{"6ba7b81": "https://practicum.yandex.ru/"},
			},
			want:       "URL with such id doesn't exist",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/"+shortURL, nil) //body)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(GetHandler(tt.input))
			handler.ServeHTTP(rr, req)

			if gotStatus := rr.Code; gotStatus != tt.wantStatus {
				t.Errorf("Want status '%d', got '%d'", tt.wantStatus, gotStatus)
			}

			// Ожидаемое (want) сообщение о ошибке должно совпадать с получаемым (got)
			if gotLocation := strings.TrimSpace(rr.Header()["Location"][0]); gotLocation != tt.want {
				t.Errorf("Want location'%s', got '%s'", tt.want, gotLocation)
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	//Здесь общие для всех тестов данные.
	shortURL := "6ba7b811"
	record := map[string]string{shortURL: "https://practicum.yandex.ru/"}

	tests := []struct {
		name       string
		ts         *storage.URLStorage
		method     string
		statusCode int
	}{
		{
			name: "all good",
			ts: &storage.URLStorage{
				Data: record,
			},
			method:     "POST",
			statusCode: http.StatusCreated,
		},
		{
			name: "bad method",
			ts: &storage.URLStorage{
				Data: record,
			},
			method: "GET",

			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/"+shortURL, nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(PostHandler(tt.ts))
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("Want status '%d', got '%d'", status, tt.statusCode)
			}
		})
	}
}
