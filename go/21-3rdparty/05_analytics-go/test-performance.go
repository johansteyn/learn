package main

import (
		"fmt"
		"os"
		"time"

		"github.com/segmentio/analytics-go"
)

func main() {
		client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
		start := time.Now()
		defer func() {
			fmt.Printf("Ended at %v...\n", time.Now())
			duration := time.Since(start)
			fmt.Printf("Took %v to send a thousand track events to Segment\n", duration)
		}()
		fmt.Printf("Starting at %v...\n", start)
		for i := 0; i < 1000; i++ {
			client.Enqueue(analytics.Track{
				Event:	fmt.Sprintf("Performance#%03d", i),
				UserId: "test-user",
			})
		}
		client.Close()
}

