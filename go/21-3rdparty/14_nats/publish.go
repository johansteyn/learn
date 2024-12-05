package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Go 3rd-party Library: NATS (publish)")
	fmt.Println()

	url := "nats://local:4T0PzzWWZoAd6jsQl9AHan2swVCHYy2w@localhost:49442"
	fmt.Printf("Connecting to NATS server using URL: %s...\n", url)
	nc, _ := nats.Connect(url)
	defer nc.Close()

	subject := "hello.world"
	message := []byte("Johan")
	fmt.Printf("Publishing message '%s' to subject '%s'...\n", message, subject)
	nc.Publish(subject, message)

	fmt.Println("Done.")

}
