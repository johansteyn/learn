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

	// Two ways to create an empty (starting) context: Background and TODO
	ctx := context.Background()
	//ctx := context.TODO()
	// They are essentially the same, but TODO serves as a temprary placeholder to
	// indicate during development that you're not yet sure what contetx will be used.
	// Do not use TODO in production code.
	// And avoid even using Background - rather use WithValue, WithCancel, etc. as the first (parent) context.

	doSomething(ctx)
	doSomethingElse(ctx, "Hello World!")

	// A context on its own is not much use.
	// A non-empty context can contain values and be cancelled.
	// Normally you pass values to functions as parameters.
	// But if you are using a framework then you may need to work with fixed function signatures.
	// For example: HTTP request handlers take request and response parameters.
	// Contexts provide an additional means of passing values to functions.

	// To add values to a context, use WithValue
	// which "wraps" the context without changing it.
	ctx = context.WithValue(ctx, "key1", "value1")
	doSomethingMore(ctx)

	// Note that using native types for keys (int, string, etc.) could potentially lead to clashes.
	// Eg: the value for key "key1" above can be overwritten:
	ctx = context.WithValue(ctx, "key1", "value2")
	doSomethingMore(ctx)
	// Rather use a custom type:
	var key1 mykey = "key1"
	ctx = context.WithValue(ctx, key1, "value3") // The value is not overridden - it's a new, separate key
	doSomethingMore(ctx)

	// Three ways to end a context: with Cancel, Deadline or Timeout
	doSomethingWithCancel(ctx)
	doSomethingWithDeadline(ctx)
	doSomethingWithTimeout(ctx)
	doSomethingWithMultipleTimeouts(ctx)

	// TODO: The AfterFunc, *Cause functions and WithoutCancel functions were added in Go 1.21

	fmt.Println("Done.")
}

func doSomething(ctx context.Context) {
	fmt.Println("Doing something with a context...")
}

// By convention, ctx is always the first parameter
// Similar to the convention of having error as last return value
func doSomethingElse(ctx context.Context, message string) {
	fmt.Printf("Doing something else with a context and message: %s\n", message)
}

type mykey string

func doSomethingMore(ctx context.Context) {
	fmt.Println("Doing something more with a context...")
	key := "key1"
	value := ctx.Value(key)
	if value != nil {
		fmt.Printf("The value for string key '%s' is '%s'\n", key, value)

	}
	var mk mykey = "key1"
	value = ctx.Value(mk)
	if value != nil {
		fmt.Printf("The value for custom key '%s' is '%s'\n", mk, value)
	}
}

func doSomethingWithCancel(ctx context.Context) {
	fmt.Println("Doing something with cancel...")
	// WithCancel returns a new context as well as a cancel function
	ctx, cancelFunc := context.WithCancel(ctx)

	go func(cancelFunc context.CancelFunc) {
		// Cancelling after 4 seconds
		// Note that the cancel function MUST be called to release resources and stop any associated goroutines
		time.Sleep(4 * time.Second)
		cancelFunc()
	}(cancelFunc)
	work(ctx)
}

// In backend (server) code you want to limit:
// - The number if concurrent requests
// - How long a request runs
// - The resources a request uses (memory, disk space, etc.)
// Go provides tolls to manage the first two: channels and contexts
// The WithTimeout and WithDeadline functions cancel a context after a specified timeout or deadline.
func doSomethingWithDeadline(ctx context.Context) {
	fmt.Println("Doing something with a deadline...")
	tenSeconds := 10 * time.Second
	deadline := time.Now().Add(tenSeconds)
	// WithDeadline also returns a cancel function, but we're ignoring it.
	ctx, _ = context.WithDeadline(ctx, deadline)
	work(ctx)
}

func doSomethingWithTimeout(ctx context.Context) {
	fmt.Println("Doing something with a timeout...")
	timeout := 8 * time.Second
	// WithTimeout is similar to WithDeadline, but takes a duration not a specific time.
	ctx, _ = context.WithTimeout(ctx, timeout)
	work(ctx)
}

func doSomethingWithMultipleTimeouts(ctx context.Context) {
	fmt.Println("Doing something with multiple timeouts...")
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	// We pass a context with a timeout of 10 seconds, but it will
	// only run for 2 seconds due to the parent context's timeout
	work(ctx)
}

func work(ctx context.Context) {
	fmt.Println("Working:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  LOOP #%d: ", i)
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			// Calling Done on a context that isn't cancellable returns nil.
			// Reading from a nil channel never returns, causing the program to hang...
			// Unless we read from it within a select statement.
			fmt.Println("cancelled!")
			return
		default:
			// The context has not been cancelled, so continue doing work...
			fmt.Println("continuing...")
		}
	}
}
