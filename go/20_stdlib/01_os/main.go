package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go Standard Library: os")
	fmt.Println()

	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("os.Executable: %s\n", executablePath)

	fileInfo, _ := os.Stdout.Stat()
	fmt.Printf("os.Stdout.Stat(): %s\n", fileInfo)
	fmt.Printf("Mode: %s\n", fileInfo.Mode())
	fmt.Printf("os.ModeCharDevice: %s\n", os.ModeCharDevice)
}

