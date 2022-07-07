package main

import (
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=yNOe3STbtGE
func main() {
	fmt.Printf("Concurrency on Youtube(1)\n")
	fmt.Println()

	// In a single thread, no refunds get processed until ALL orders have been processed
	//process("order", 1_000_000, 500)
	//process("refund", 6, 200)
	//fmt.Println("Done")

	// Have orders and refunds processed concurrently,
	// but it is "done" before a single job gets processed...
	//go process("order", 1_000_000, 500)
	//go process("refund", 6, 200)
	//fmt.Println("Done")

	// A naive way of blocking...
	//go process("order", 1_000_000, 500)
	//go process("refund", 6, 200)
	//// Wait for a return key press 
	//fmt.Scanln()
	//fmt.Println("Done")

	// We need a way for the processing thread to notify the main thread that it's done.
	// But we want to do that without modifying the process function.
	// We can do this with a Wait Group
	/*
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			process("order", 10, 500)
			wg.Done()
		}()
		go func() {
			process("refund", 12, 200)
			wg.Done()
		}()
		fmt.Println("Waiting...")
		wg.Wait()
		fmt.Println("Done")
	*/

	// But there's something better than wait groups: channels
	/*
		channel := make(chan string)
		go func() {
			process("order", 10, 500)
			channel <- "Done processing orders"
		}()
		go func() {
			process("refund", 12, 200)
			channel <- "Done processing refunds"
		}()
		// We can wait until we have received both messages from the channel
		message := <-channel
		fmt.Printf("Received channel message: %s\n", message)
		message = <-channel
		fmt.Printf("Received channel message: %s\n", message)
		// We can also loop through the channel, but then we need to know how many times to loop,
		// or we need to close the channel when we know we're done...
		//for message := range channel {
		//	fmt.Printf("Received channel message: %s\n", message)
		//}
	*/

	// We can use separate channels
	orderChannel := make(chan string)
	go func() {
		process("order", 10, 500)
		orderChannel <- "Done processing orders"
	}()
	refundChannel := make(chan string)
	go func() {
		process("refund", 12, 200)
		refundChannel <- "Done processing refunds"
	}()
	// We can wait until we have received both messages from the channel
	//message := <-orderChannel
	//fmt.Printf("Received channel message: %s\n", message)
	//message = <-refundChannel
	//fmt.Printf("Received channel message: %s\n", message)
	// But refunds are processed faster than orders, yet we need to wait for orders to complete...
	// Instead, we can use "select":
	var ordersDone, refundsDone bool
	for !ordersDone || !refundsDone {
		select {
		case message := <-orderChannel:
			fmt.Printf("Received channel message: %s\n", message)
			ordersDone = true
		case message := <-refundChannel:
			fmt.Printf("Received channel message: %s\n", message)
			refundsDone = true
		}
	}

}

func process(job string, number int, duration time.Duration) {
	for i := 0; i < number; i++ {
		fmt.Printf("Processing %s #%d\n", job, i)
		pause(duration)

	}
}

func pause(millis time.Duration) {
	time.Sleep(millis * time.Millisecond)
}
