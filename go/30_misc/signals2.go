package main

//https://golangcode.com/handle-ctrl-c-exit-in-terminal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Signals")
	setupSignalHandler()
	fmt.Print("Awaiting signal")
	for {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("Done.")
}

func setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,  // CTRL-C
		syscall.SIGTSTP, // CTRL-Z
		syscall.SIGQUIT, // CTRL-\
		syscall.SIGTERM) // kill command
	go func() {
		sig := <-c
		fmt.Println()
		fmt.Printf("Signal intercepted: %s\n", sig)
		cleanup()
		os.Exit(1)
	}()
}

func cleanup() {
	fmt.Println("Cleaning up...")
}
