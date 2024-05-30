package main

import (
	"context"
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
	fmt.Println()

	run1()
	fmt.Println("Completed run1.")
	fmt.Println()

	// Need to pause because we're waiting in a seperate thread, otherwise we get:
	//   panic: sync: WaitGroup is reused before previous Wait has returned
	pause(4)

	run2()
	fmt.Println("Completed run2.")
	fmt.Println()

	pause(4)

	fmt.Println("Done.")
}

var count int
var wg sync.WaitGroup
var shutdownRequested bool
var cancelFunc context.CancelFunc

// A simple way to detect shutdown using a boolean flag
func run1() {
	listenForInterrupt(shutdown1)
	for count = 1; count <= 10; count++ {
		if shutdownRequested {
			fmt.Printf("run1: shutdown requested during loop #%d, returning...\n", count)
			return
		}
		fmt.Printf("run1: starting job #%d...\n", count)
		work(count)
	}
}

// A better way to detect shutdown is to use a context
func run2() {
	ctx, c := context.WithCancel(context.Background())
	cancelFunc = c
	listenForInterrupt(shutdown2)
	for count = 1; count <= 10; count++ {
		select {
		case <-ctx.Done():
			fmt.Printf("run2: context cancelled during loop #%d, returning...\n", count)
			return
		default:
			fmt.Printf("run2: starting job #%d...\n", count)
			work(count)
		}
	}

}

func work(i int) {
	wg.Add(1)
	defer wg.Done()
	fmt.Printf("[#%d]Working...\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("[#%d]Done.\n", i)
}

// Allow graceful shutdown when the user presses Ctrl+C
func listenForInterrupt(shutdown func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received signal...")
		shutdown()
	}()
}

func shutdown1() {
	fmt.Println("shutdown1: setting flag...")
	shutdownRequested = true
	fmt.Printf("shutdown1: waiting for job #%d to finish...\n", count)
	wg.Wait()
	fmt.Println("Done waiting.")
}

func shutdown2() {
	fmt.Println("shutdown2: cancelling...")
	cancelFunc()
	fmt.Printf("shutdown2: waiting for job #%d to finish...\n", count)
	wg.Wait()
	fmt.Println("Done waiting.")
}

func pause(seconds int) {
	fmt.Printf("Sleeping for %d seconds...\n", seconds)
	time.Sleep(time.Second * time.Duration(seconds))
}
