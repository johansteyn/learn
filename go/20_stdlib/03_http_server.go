package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func batch(w http.ResponseWriter, req *http.Request) {
	fmt.Println("batch")
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "batch\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/v1/batch", batch)
	http.ListenAndServe(":8080", nil)
}

