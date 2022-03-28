package main

import (
	"fmt"
	"net"
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

	fmt.Printf("Sending datagram to: %s...\n", c.RemoteAddr().String())
	data := []byte("Hello World!\n")
	_, err = c.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done.")
}

