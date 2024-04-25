package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "http://localhost:8080"
	args := os.Args[1:]
	if len(args) >= 1 {
		url = "http://" + args[0]
	}
	c := http.Client{
		Timeout: 1 * time.Second,
	}
	var start, end time.Time
	defer func() {
	  duration := end.Sub(start)
		fmt.Printf("Time taken: %v\n", duration)
	}()
	fmt.Printf("Sending HTTP request to %s...\n", url)
	start = time.Now()
	_, err := c.Get(url)
	end = time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}
}

