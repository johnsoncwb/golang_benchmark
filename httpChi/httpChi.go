package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{name}", ChiHelloWorld)
	http.ListenAndServe("localhost:9000", r)
}

func ChiHelloWorld(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Olá %s!!!", name),
	})
}
