package main

import (
	"context"
	"fmt"
	"time"
)

// https://pkg.go.dev/context
// https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
// https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea
func main() {
	fmt.Println("Go Standard Library: context")
	fmt.Println()

	// Two ways to create an empty (starting) context: TODO and Background
	// They are essentially the same, but using Background indicates intent to other developers
	// (whatever that means...)
	//ctx := context.TODO()
	ctx := context.Background()

	doSomething(ctx)
	doSomethingElse(ctx, "Hello World!")

	// A context on its own is not much use.
	// To add values to a context, use WithValue
	// which "wraps" the context without changing it.
	ctx = context.WithValue(ctx, "key1", "value1")
	doSomethingMore(ctx)

	doSomethingWithCancel(ctx)

	doSomethingWithDeadline(ctx)

	fmt.Println("Done.")
}

func doSomething(ctx context.Context) {
	fmt.Println("Doing something with a context...")
}

// By convention, ctx is always the first parameter
func doSomethingElse(ctx context.Context, message string) {
	fmt.Printf("Doing something else with a context and message: %s\n", message)
}

func doSomethingMore(ctx context.Context) {
	fmt.Println("Doing something more with a context...")
	key := "key1"
	value := ctx.Value(key)
	fmt.Printf("The value for key '%s' is '%s'\n", key, value)
}

func doSomethingWithCancel(ctx context.Context) {
	fmt.Println("Doing something with cancel...")
	// WithCancel returns a new context as well as a cancel function
	ctx, cancelFunc := context.WithCancel(ctx)

	go func(cancelFunc context.CancelFunc) {
		// Cancelling after 4 seconds
		time.Sleep(4 * time.Second)
		cancelFunc()
	}(cancelFunc)

	work(ctx)
}

func doSomethingWithDeadline(ctx context.Context) {
	fmt.Println("Doing something within a deadline...")
	tenSeconds := 10 * time.Second
	deadline := time.Now().Add(tenSeconds)
	// WithDeadline also returns a cancel function, but we're ignoring it.
	ctx, _ = context.WithDeadline(ctx, deadline)

	work(ctx)
}

func work(ctx context.Context) {
	fmt.Println("Doing some work...")
	for i := 0; i < 10; i++ {
		fmt.Printf("  LOOP #%d: ", i)
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("cancelled!")
			return
		default:
			// The context has not been cancelled, so continue doing work...
			fmt.Println("continuing...")
		}
	}
}
