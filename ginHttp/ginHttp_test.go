package main

import (
	"github.com/gin-gonic/gin"
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

		router := gin.New()
		gin.SetMode(gin.ReleaseMode)
		router.GET("/:name", GinHelloWorld)
		router.ServeHTTP(rr, req)
	}
}
