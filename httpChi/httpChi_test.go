package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkChi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		req, err := http.NewRequest("GET", "/Allan", nil)
		if err != nil {
			b.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := chi.NewRouter()
		router.Get("/{name}", ChiHelloWorld)
		router.ServeHTTP(rr, req)
	}
}
