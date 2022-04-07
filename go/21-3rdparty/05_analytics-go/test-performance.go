package main

import (
	"fmt"
	"net"
	"net/http"
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
//	client := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))

	// Try to force HTTP1 instea dof HTTP2
	// https://stackoverflow.com/questions/67770829/go-http-request-falls-back-to-http2-even-when-force-attempt-is-set-to-false
  //os.Setenv("GODEBUG", "http2client=0")
	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
//			Timeout:   30 * time.Second,
//			Timeout:   300 * time.Millisecond,
			Timeout:   time.Second,
//			KeepAlive: 30 * time.Second,
			KeepAlive: time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
//		ForceAttemptHTTP2:     false,
		MaxIdleConns:          100,
//		IdleConnTimeout:       90 * time.Second,
		IdleConnTimeout:       time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client, _ := analytics.NewWithConfig(
		os.Getenv("SEGMENT_WRITE_KEY"),
		analytics.Config{
			// Default endpoint is https://api.segment.io
			// Deliberately incorrect to force timeout
			//Endpoint: "https://segment.io",
			//Endpoint: "https://johansteyn.eu",
			//Endpoint: "http://localhost:8080",
//			Transport: transport,
			BatchSize: 1, // Default batch size is 250
	})
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

