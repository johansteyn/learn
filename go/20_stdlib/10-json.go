package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// https://pkg.go.dev/encoding/json
func main() {
	fmt.Println("Go Standard Library: json")
	fmt.Println()

	s := "The quick brown fox"
	bytes, err := json.Marshal(s)
	if err != nil {
		handleError("Error marshalling string", err)
	}
	fmt.Printf("Marshalled bytes: %v\n", bytes)
	fmt.Printf(" Bytes as string: %s\n", bytes)

	fmt.Println("Done.")
}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
