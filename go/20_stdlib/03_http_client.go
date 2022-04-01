package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout: 1 * time.Second,
	}
	var start, end time.Time
	defer func() {
	  duration := end.Sub(start)
		fmt.Printf("Time taken: %v\n", duration)
	}()
	url := "http://localhost:8080/hello"
	fmt.Printf("Sending HTTP request to %s...\n", url)
	start = time.Now()
	_, err := c.Get(url)
	end = time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}
}

