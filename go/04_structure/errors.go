package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Errors")
	fmt.Println()

	// Invalid values
	//x := -123
	//x := 999
	//y := 0
	//y := -1
	//y := 12
	//y := 3
	// Valid values
	x := 9
	y := 4

	// Return values, including the error value, should all be read
	// Though it's not mandatory to read return values...
	div(x, y) // Works - don't need to read return values

	// If you read any return values then you need to read all of them
	//result, remainder := div(x, y) // Doesn't work - need to read all or nothing

	// A return value can be ignored with underscore
	// But it's not good practice to ignore errors!
	// Except for functions like fmt.Println
	//result, remainder, _ := div(x, y)

	// Best practice is to read all return values, including the error
	result, remainder, err := div(x, y)
	if err != nil {
		// Passing an error as parameter to Println calls the Error method automatically
		fmt.Println(err)
		// But it's more common to use Printf
		fmt.Printf("Error: %v\n", err)
		var e MyOtherError
		if errors.As(err, &e) {
			fmt.Printf("Code: %d\n", e.Code)
		}
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Printf("Wrapped error: %v\n", wrappedErr)
		}
		os.Exit(1)
	}
	fmt.Println("No error 🙂")
	fmt.Printf("div(%d, %d): result=%d, remainder=%d\n", x, y, result, remainder)

	fmt.Println()
}

// Allow only numerators between -100 and 100, and only non-zero and positive even denominators smaller than 10
func div(numerator, denominator int) (int, int, error) {
	// By convention, the last return value is of type error
	if numerator < -100 {
		// Using errors.New to create an error
		return 0, 0, errors.New("numerator is too small")
	}
	if numerator > 100 {
		// Using fmt.Errorf to create an error
		return 0, 0, fmt.Errorf("numerator is too large: %d", numerator)
	}
	if denominator == 0 {
		// Using fmt.Errorf to wrap an error
		e := errors.New("wrapped error")
		return 0, 0, fmt.Errorf("denominator is zero: %d, root cause: %w", denominator, e)
	}
	if denominator < 0 {
		// Using my custom error
		return 0, 0, MyError("denominator is negative")
	}
	if denominator >= 10 {
		// Using my other custom error, with an error code
		return 0, 0, MyOtherError{99, "denominator is too large"}
	}
	if denominator%2 == 1 {
		// Using a custom wrapping error
		e := MyError("wrapped error")
		return 0, 0, MyWrappingError{"denominator is not even", e}
	}
	// Note that error messages must not be capitalized nor end with punctuation or a newline
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}

// A custom error
type MyError string

// This satisfies the "error" interface,
// which consists of a single "Error" function
// that takes no parameters and returns a string
func (e MyError) Error() string {
	return string(e)
}

// Another custom error
type MyOtherError struct {
	Code    int
	Message string
}

// This also satisfies the "error" interface,
func (e MyOtherError) Error() string {
	return e.Message
}

// A wrapping  error
type MyWrappingError struct {
	message string
	cause   error
}

func (e MyWrappingError) Error() string {
	return e.message
}

func (e MyWrappingError) Unwrap() error {
	return e.cause
}

// TODO:
// https://www.youtube.com/watch?v=0c-1KJwSMCw
// @8:30
