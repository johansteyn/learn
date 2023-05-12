package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Bio struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("gorilla/mux")
	fmt.Println()

	router := mux.NewRouter()
	fmt.Printf("*** router: %v\n", router)

	router.Path("/helloworld").Methods(http.MethodGet).HandlerFunc(HelloWorldHandler)
	router.Path("/outer").Methods(http.MethodGet).HandlerFunc(OuterHandler)
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("*** Books handler")
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You requested page %s of book %s\n", page, title)
	})

	// http://localhost/helloworld
	// http://localhost/books/aaa/page/1
	// Somehow this works with port 80, even when not running as root...
	http.ListenAndServe(":80", router)
}

func HelloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("*** HelloWorldHandler")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Hello World!")
}

func OuterHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("*** OuterHandler")
	HelloWorldHandler(rw, r)
}
