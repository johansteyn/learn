package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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

	// Allow graceful shutdown when the user presses Ctrl+C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	count = 0
	go func() {
		<-c
		fmt.Printf("Received signal while doing job #%d...\n", count)
		fmt.Println("Calling shutdown...")
		shutdown()
	}()
	for count = 1; count <= 10; count++ {
		if shutdownRequested {
			fmt.Println("Breaking...")
			break
		}
		fmt.Printf("Starting job #%d...\n", count)
		work(count)
	}
	fmt.Println("Sleeping for 4 seconds...")
	time.Sleep(4 * time.Second)
	fmt.Println("Completed without interruption.")
}

var count int
var wg sync.WaitGroup

var shutdownRequested bool

func work(i int) {
	wg.Add(1)
	defer wg.Done()
	fmt.Printf("[#%d]Working...\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("[#%d]Done.\n", i)
}

func shutdown() {
	shutdownRequested = true
	fmt.Printf("Shutdown, waiting for job #%d to finish...\n", count)
	wg.Wait()
	// No guarantee that anything will be called after waiting.
	// But the 4-second sleep in main gives enough time...
	fmt.Println("Exiting...")
	os.Exit(0)
}
