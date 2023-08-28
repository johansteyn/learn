package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	// Using this before the logger inserts a unique ID for each request
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	// A recoverer handles panics by logging the panic (with backtrace) and returning 500 (Internal Server Error) response.
	router.Use(middleware.Recoverer)
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("root"))
	})
	router.Get("/helloworld", HelloWorldHandler)
	router.Get("/panic", PanicHandler)
	http.ListenAndServe(":80", router)
}

func HelloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("*** HelloWorldHandler")
	fmt.Printf("=== Remote Address: %v\n", r.RemoteAddr)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Hello World!")
}

func PanicHandler(rw http.ResponseWriter, r *http.Request) {
	var zero = 0
	fmt.Printf("=== Divide by zero: %d\n", 1/zero)
	// This statement is not reached...
	rw.Write([]byte("PANIC!"))
}
