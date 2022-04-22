package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/segmentio/analytics-go"
)

// TODO: Duplicate - move to library...
type Event struct {
	Timestamp string `json:"timestamp"`
	Source string `json:"source"`
	Name string `json:"name"`
	AnonymousId string `json:"anonymousId"`
	UserId string `json:"userId"`
	Properties map[string]string `json:"properties"`
}

func main() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	segmentClient := analytics.New(os.Getenv("SEGMENT_WRITE_KEY"))
	defer segmentClient.Close()

	// TODO: Is 1024 bytes enough?
	// Maybe for a single event, but certainly not for a list of events...
	// Even a single event might not fit, eg: if error message is long.
	// So, client should only send a single event at a time,
	// and ensure each message is no longer than 1024 bytes.
	buffer := make([]byte, 1024)
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		message := string(buffer[0:n-1])
		fmt.Printf("\nReceived from %v: %s\n", addr, message)
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}
		event := Event{}
		json.Unmarshal(buffer[0:n-1], &event)
		fmt.Printf("Event: %#v\n", event)

		properties := analytics.NewProperties()
		for k, v := range event.Properties {
			properties.Set(k, v)
		}

		segmentClient.Enqueue(analytics.Track{
			Event: event.Name,
			AnonymousId: event.AnonymousId,
			UserId: event.UserId,
			Properties: properties,
		})
	}
}

