package handlers

import (
	"net/http"
	"reflect"
	"testing"
)

func TestPostHandler(t *testing.T) {
	type args struct {
		urlSaver URLSaver
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostHandler(tt.args.urlSaver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	type args struct {
		urlGeter URLGeter
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHandler(tt.args.urlGeter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
