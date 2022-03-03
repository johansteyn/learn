package main

import (
	"os"

	"github.com/segmentio/analytics-go"
)

func main() {
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	properties := analytics.NewProperties()
	properties.Set("command", "mongocli iam projects ls")
	properties.Set("duration", "123ms")
	properties.Set("result", "SUCCESS")
	client.Enqueue(analytics.Track{
		Event: "test-properties",
		UserId: "Alice",
		Properties: properties,
	})
	client.Close()
}

