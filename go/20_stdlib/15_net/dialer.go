package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	url := "http://localhost:8080"
	args := os.Args[1:]
	if len(args) >= 1 {
		url = "http://" + args[0]
	}
	roundTripper := New(true)
	c := http.Client{
		Transport: roundTripper,
	}
	fmt.Printf("Sending HTTP request to %s...\n", url)
	_, err := c.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// func New(rejectLocalConnections bool) http.RoundTripper {
func New(rejectLocalConnections bool) *http.Transport {
	dialer := NewDialer(rejectLocalConnections)
	return &http.Transport{
		DialContext: dialer.DialContext,
	}
}

func NewDialer(rejectLocalConnections bool) *net.Dialer {
	var controlFn func(network, address string, c syscall.RawConn) error
	if rejectLocalConnections {
		controlFn = RejectLocalConnections
	}
	return &net.Dialer{
		Control: controlFn,
	}
}

func RejectLocalConnections(network, address string, c syscall.RawConn) error {
	fmt.Printf("Checking if address is local: %s\n", address)
	ip, _, err := net.SplitHostPort(address)
	if err != nil {
		return errors.New("invalid address")
	}
	if IP := net.ParseIP(ip); IP.IsPrivate() || IP.IsMulticast() || IP.IsUnspecified() || IP.IsLoopback() {
		return errors.New("address not allowed")
	}
	fmt.Println("Address is not local :)")
	return nil
}
