package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	natsServer "github.com/nats-io/nats-server/v2/server"
	nats "github.com/nats-io/nats.go"
)

// This is for if you want to embed a NATS server in your application:
//
//	https://www.youtube.com/watch?v=cdTrl8UfcBo
//	https://medium.com/cloud-native-daily/create-real-time-pub-sub-with-nats-in-go-5f88d13c4927
//
// Otherwise run the NATS server using the NATS CLI:
//
//	$ nats server run
//
// Then use the user credentials output to stdout
func main() {
	fmt.Println("Go 3rd-party Library: NATS (embedded)")
	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
		<-sigChannel
		close(sigChannel)
		cancel()
	}()

	createServer()

	go consumer(ctx)

	go producer(ctx)
	<-ctx.Done()

	fmt.Println("Done.")
}

func createServer() {
	opts := &natsServer.Options{
		Host:   "localhost",
		Port:   49443,
		NoLog:  false,
		NoSigs: false,
	}

	fmt.Println("Creating embedded NATS server...")
	server, err := natsServer.NewServer(opts)
	if err != nil {
		handleError("Error creating NATS server.", err)
	}

	fmt.Println("Starting...")
	err = natsServer.Run(server)
	if err != nil {
		handleError("Error starting NATS server.", err)
	}

	fmt.Println("Running!")
}

func consumer(ctx context.Context) {
	fmt.Println("[consumer] Connecting to the NATS server...")
	connection, err := nats.Connect("nats://localhost:49443")
	if err != nil {
		handleError("[consumer] Error connecting to NATS server.", err)
	}
	defer connection.Close()

	subject := "hello.world"

	fmt.Printf("[consumer] Subscribing to subject '%s'...\n", subject)
	messages := make(chan *nats.Msg, 1000)
	subscription, err := connection.ChanSubscribe(subject, messages)
	if err != nil {
		handleError("[consumer] Error subscribing to subject.", err)
	}

	defer func() {
		subscription.Unsubscribe()
		close(messages)
	}()

	fmt.Println("[consumer] Listening for messages...")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[consumer] Finished consuming messages.")
			return
		case message := <-messages:
			fmt.Printf("[consumer] Received message: %s\n", message.Data)
		}
	}
}

func producer(ctx context.Context) {
	fmt.Println("[producer] Connecting to the NATS server...")
	// The embedded server
	url := "nats://localhost:49443"
	// External NATS CLI server
	//url := "nats://local:4T0PzzWWZoAd6jsQl9AHan2swVCHYy2w@localhost:49442"
	connection, err := nats.Connect(url)
	if err != nil {
		handleError("[producer] Error connecting to NATS server.", err)
	}
	defer connection.Close()

	subject := "hello.world"
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[producer] Finished publishing messages.")
			return
		default:
			i++
			message := []byte(fmt.Sprintf("Message %d", i))
			fmt.Printf("[producer] Publishing message '%s' to subject '%s'...\n", message, subject)
			connection.Publish(subject, message)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
