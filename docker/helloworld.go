package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("helloworld")
	fmt.Fprintf(w, "Hello World!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


func main() {
	fmt.Println("Starting 'helloworld' server...")
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Done.")
}

