package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Standard Library: time")
	fmt.Println()

	//layout := "2006-01-02T15:04:05.999Z"
	//layout := "2006-01-02T15:04:05.000Z"
	layout := "2006-01-02T15:04:05.000Z"
	now := time.Now()
	fmt.Printf("Now: %s\n", now.UTC().Format(layout))

	/*
		//then, err := time.Parse(layout, "2020-01-01T00:00:00.000Z")
		//then, err := time.Parse(layout, "2025-01-09T12:44:15.24Z")
		then, err := time.Parse(layout, "2025-01-09T12:44:15.240Z")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Then: %s\n", then.UTC().Format(layout))
	*/
	fmt.Println()
}
