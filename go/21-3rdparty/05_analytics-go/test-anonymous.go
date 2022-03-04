package main

import (
	"fmt"
	"os"

	"github.com/segmentio/analytics-go"
	"github.com/segmentio/ksuid"
)

func main() {
	id := ksuid.New()
	fmt.Printf("Generated anonymous ID: %v\n", id)
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	client.Enqueue(analytics.Track{
		Event: "test-anonymous",
		AnonymousId: id.String(),
		// Added the UserId to see if both can be used...
		UserId: "test-user",
	})
	client.Close()
}

