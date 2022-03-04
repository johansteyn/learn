package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/analytics-go"
)

func main() {
	args := os.Args[1:]
	numEvents := 1000
	prefix := "p"
	if len(args) >= 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid number: %s\n", args[0])
			os.Exit(1)
		}
		numEvents = i
	}
	if len(args) == 2 {
		prefix = args[1]
	}
	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	start := time.Now()
	defer func() {
		fmt.Printf("Ended at %v...\n", time.Now())
		duration := time.Since(start)
		fmt.Printf("Took %v to send %d track events\n", duration, numEvents)
	}()
	fmt.Printf("Starting at %v...\n", start)
	for i := 0; i < numEvents; i++ {
		client.Enqueue(analytics.Track{
			Event:	fmt.Sprintf("%s-%06d", prefix, i),
			UserId: "test-user",
		})
	}
	client.Close()
}

