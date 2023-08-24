package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Root path"))
	})
	router.Get("/helloworld", HelloWorldHandler)
	http.ListenAndServe(":80", router)
}

func HelloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("*** HelloWorldHandler")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Hello World!")
}
