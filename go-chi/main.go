package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("starting server")
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Compress(6, "gzip"))
	r.Use(middleware.StripSlashes)

	r.Post("/chi-test", apiFunc)

	http.ListenAndServe(":3000", r)

}

func apiFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	w.Write([]byte("{}"))
	w.WriteHeader(http.StatusAccepted)
}
