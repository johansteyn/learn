package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", "localhost:8080")
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

	message := "Hello World!"
	args := os.Args[1:]
	if len(args) > 0 {
		message = args[0]
	}

	var start, end time.Time
	fmt.Printf("Sending datagram to %s...\n", c.RemoteAddr().String())
	defer func() {
	  duration := end.Sub(start)
		fmt.Printf("Time taken: %v\n", duration)
	}()
	data := []byte(message + "\n")
	start = time.Now()
	_, err = c.Write(data)
	end = time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}
}
