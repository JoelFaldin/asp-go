package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("greetings from chi backend"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", test)

	http.ListenAndServe(":3000", r)
}
