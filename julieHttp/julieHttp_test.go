package main

import (
	"github.com/julienschmidt/httprouter"
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

		router := httprouter.New()
		router.GET("/{name}", JulieHelloWorld)
		router.ServeHTTP(rr, req)
	}
}
