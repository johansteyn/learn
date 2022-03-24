package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Signals")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,
		syscall.SIGINT,  // CTRL-C
		syscall.SIGTSTP, // CTRL-Z
		syscall.SIGQUIT, // CTRL-\
		syscall.SIGTERM) // kill command
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("Signal intercepted: %s\n", sig)
		done <- true
	}()
	fmt.Println("Awaiting signal...")
	<-done
	fmt.Println("Done.")
}

