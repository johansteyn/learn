package main

import (
    "os"

    "github.com/segmentio/analytics-go"
)

func main() {
    // Instantiates a client to use send messages to the segment API.
    client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))

    // Enqueues a track event that will be sent asynchronously.
    client.Enqueue(analytics.Track{
        UserId: "test-user",
        Event:  "test-snippet",
    })

    // Flushes any queued messages and closes the client.
    client.Close()
}

