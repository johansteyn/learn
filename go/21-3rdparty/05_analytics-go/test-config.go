package main

import (
	"os"
	"time"

	"github.com/segmentio/analytics-go"
)

func main() {
	client, _ := analytics.NewWithConfig(
		os.Getenv("SEGMENT_WRITE_KEY"),
		analytics.Config{
			// If not specified, uses the default of https://api.segment.io
			Endpoint: "https://white.10gen.io",
			Interval:  30 * time.Second,
			BatchSize: 100,
			Verbose:   true,
	})
	defer client.Close()

	client.Enqueue(analytics.Track{
		Event: "test-config",
		UserId: "test-user",
	})

}

