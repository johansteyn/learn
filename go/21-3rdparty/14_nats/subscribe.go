package main

import (
	"fmt"
	"os"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Go 3rd-party Library: NATS (subscribe)")
	fmt.Println()

	url := "nats://local:4T0PzzWWZoAd6jsQl9AHan2swVCHYy2w@localhost:49442"
	fmt.Printf("Connecting to NATS server using URL: %s...\n", url)
	nc, err := nats.Connect(url)
	if err != nil {
		handleError("Error connecting to NATS server.", err)
	}
	defer nc.Close()

	subject := "hello.world"
	//message := []byte("Johan")
	//fmt.Printf("Publishing message '%s' to subject '%s'...\n", message, subject)
	//nc.Publish(subject, message)

	fmt.Printf("Subscribing (synchronously) to subject '%s'...\n", subject)
	subscription, err := nc.SubscribeSync(subject)
	if err != nil {
		handleError("Error subscribing to subject.", err)
	}
	//msg, err := subscription.NextMsg(10 * time.Millisecond)
	//msg, err := subscription.NextMsg(10 * time.Second)
	msg, err := subscription.NextMsg(time.Minute)
	if err != nil {
		handleError("Error receiving message.", err)
	}
	fmt.Printf("Received message: %s\n", msg.Data)
	fmt.Println()

	fmt.Printf("Subscribing (asynchronously) to subject '%s'...\n", subject)
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", m.Data)
	})
	if err != nil {
		handleError("Error subscribing to subject.", err)
	}

	fmt.Println("Done.")

}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
