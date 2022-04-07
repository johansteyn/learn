package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type Event struct {
	Timestamp string `json:"timestamp"`
	Source string `json:"source"`
	Name string `json:"name"`
	AnonymousId string `json:"anonymousId"` // Hopefully won't need this...
	UserId string `json:"userId"`
	ProfileId string `json:"profileId"` // Either "default" or a hash of the profile name
	OrgId string `json:"orgId"`
	ProjectId string `json:"projectId"`
	Service string `json:"service"`
	Authentication string `json:"authentication"`
	Duration string `json:"duration"`
	Result string `json:"result"`
	Error string `json:"error"`
	Flags string `json:"flags"`
	Alias string `json:"alias"`
	Version string `json:"version"`
	OS string `json:"os"`
	Installer string `json:"installer"`
	Terminal string `json:"terminal"`
}

// A new event will be created with the current time as timestamp.
// If no address is specified, then the event will be persisted to file.
// If an address is specified, then any previously persisted events
// as well as the new event events will be sent to that address.
// Eg: localhost:8080
func main() {
	var addr string
	args := os.Args[1:]
	if len(args) >= 1 {
		addr = args[0]
	}

	now := time.Now()
	// Experimenting with different time format strings...
	//fmt.Printf("now: %v (%T)\n", now, now)
	//fmt.Printf("now.String(): %s\n", now.String())
	//fmt.Printf("now.Format(time.RFC3339): %s\n", now.Format(time.RFC3339))
	//fmt.Printf("now.Format(time.RFC3339Nano): %s\n", now.Format(time.RFC3339Nano))
	//RFC3339Millis := "2006-01-02T15:04:05.999Z07:00"
	//fmt.Printf("now.Format(custom): %s\n", now.Format(RFC3339Millis))

	event := &Event{
		Timestamp: now.Format(time.RFC3339Nano),
		Source: "atlascli",
		Name: "atlas-projects-list",
		UserId: "123abc456def789ghi987jkl654mno321pqr123stu456vwx",
		ProfileId: "default",
		OrgId: "1234567890abcdefghijklmn",
		ProjectId: "0987654321zyxwvutsrqponm",
		Service: "cloud",
		Authentication: "OAuth",
		Duration: "123",
		Result: "SUCCESS",
		Flags: "--limit,--page",
		Alias: "ls",
		Version: "1.23.0",
		OS: "linux/amd64",
		Installer: "homebrew",
		Terminal: "teletype",
	}

	filename := "events.json"
	if len(addr) == 0 {
		fmt.Printf("No UDP address specified, so persisting event in file: %s\n", filename)
		err := writeEvent(filename, event)
		if err != nil {
			fmt.Println("Error writing JSON file")
			os.Exit(1)
		}
		return
	}

	s, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	events, err := readEvents(filename)
	if err != nil {
			fmt.Println("Error reading JSON file")
			os.Exit(1)
	}
	fmt.Printf("Events: %#v\n", events)

	events = append(events, *event)
	for _, event := range events {
		data, _ := json.Marshal(event)
		message := string(data)
		fmt.Printf("Message:\n %s\n", message)
		var start, end time.Time
		fmt.Printf("Sending datagram to %s...\n", c.RemoteAddr().String())
		datagram := []byte(message + "\n")
		start = time.Now()
		_, err = c.Write(datagram)
		end = time.Now()
		if err != nil {
			fmt.Println(err)
			return
		}
		duration := end.Sub(start)
		fmt.Printf("Time taken: %v\n\n", duration)
	}

	_, err = os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			fmt.Println("Error deleting JSON file")
			os.Exit(1)
		}
	}
}

// Read persisted events from file, if it exists
func readEvents(filename string) ([]Event, error) {
	var events = []Event{}
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return events, nil
	}
	file, err := os.Open("events.json")
	if err != nil {
			return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var event Event
		decoder.Decode(&event) // Reads the next JSON-encoded value
		events = append(events, event)
	}
	return events, nil
}

// https://dev.to/evilcel3ri/append-data-to-json-in-go-5gbj
func writeEvent(filename string, event *Event) (error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
			if err != nil {
				return err
			}
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, _ := json.MarshalIndent(event, "", "\t")
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

