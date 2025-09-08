package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Errors")
	fmt.Println()

	x := 0 // No error
	//x := 1 // errors.New
	//x := 2 // fmt.Errorf
	//x := 3 // Sentinel error
	//x := 4 // MyError
	//x := 5 // MyOtherError
	//x := 6 // errors.New wrapped in fmt.Errorf
	//x := 7 // MyError wrapped in MyWrappingError
	//x := 8 // Sentinel error wrapped in MyWrappingError
	//x := 9 // Two levels of wrapping
	//x := 10 // Multiple errors
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid value specified:", os.Args[1])
			os.Exit(1)
		}
		x = i
	}
	fmt.Printf("x = %d\n", x)
	fmt.Println()

	// Return values, including the error value, should all be read
	// Though it's not mandatory to read return values...
	increment(x) // Works - don't need to read return values

	// If you read any return values then you need to read all of them
	//result, remainder := div(x, y) // Doesn't work - need to read all or nothing

	// A return value can be ignored with underscore
	// But it's not good practice to ignore errors!
	// Except for functions like fmt.Println
	//result, remainder, _ := div(x, y)

	// Best practice is to read all return values, including the error
	result, err := increment(x)
	if err != nil {
		// Passing an error as parameter to Println calls the Error method automatically
		fmt.Println("Error (println):")
		fmt.Println(err)
		fmt.Println()

		// But it's more common to use Printf
		fmt.Printf("Error (printf): %v\n", err)
		fmt.Println()

		fmt.Printf("Error type: %T\n", err)
		fmt.Println()

		// Check if the error wraps another error
		wrappedErr := errors.Unwrap(err)
		fmt.Printf("Wrapped error: %v\n", wrappedErr)
		fmt.Println()

		// Check if the error wraps multiple (joined) errors
		// NOTE: For some reason errors that are joined don't unwrap...
		// https://www.reddit.com/r/golang/comments/15x9rm0/why_there_is_no_unwrapjoin_in_errors_package/
		// https://github.com/golang/go/issues/69586
		fmt.Println("Wrapped errors:")
		for err := errors.Unwrap(err); err != nil; err = errors.Unwrap(err) {
			fmt.Printf("  %v\n", err)
		}
		fmt.Println()

		// Check if the err is the sentinel one
		if err == ErrSentinel {
			fmt.Println("Error == the sentinel error")
		} else {
			fmt.Println("Error != the sentinel error")
		}
		fmt.Println()

		// Use errors.Is to check if the error is or wraps the specified error
		if errors.Is(err, ErrSentinel) {
			fmt.Println("Error is the sentinel error")
		} else {
			fmt.Println("Error is NOT the sentinel error")
		}
		fmt.Println()

		// Use errors.As to check if the error or any error it wraps is of the specified type
		// ie. if it is assignable to the specified error
		// But be careful when using errors.As to check against a type - not an actual value
		// Eg: Don't check against the sentinel error as it can change its value
		//fmt.Printf("SentinelError before calling errors.As: %v\n", ErrSentinel)
		//if errors.As(err, &ErrSentinel) {
		//	fmt.Println("Error as the sentinel error")
		//}
		//fmt.Printf("SentinelError after calling errors.As: %v\n", ErrSentinel)
		//fmt.Println()

		// Also, it can result in a panic if the second parameter is not of the correct type
		//if errors.As(err, fs.ErrNotExist) {
		//	fmt.Println("Error is of type fs.ErrNotExist")
		//}

		// The correct way to use errors.As is to check against a type
		var me MyError
		if errors.As(err, &me) {
			fmt.Println("Error as MyError")
		} else {
			fmt.Println("Error as NOT MyError")
		}
		var moe MyOtherError
		if errors.As(err, &moe) {
			// Since errors.As can assign the value of the error to the variable, we can access the fields
			fmt.Printf("Error as MyOtherError. Code: %d\n", moe.Code)
		} else {
			fmt.Println("Error as NOT MyOtherError")
		}
		var mwe MyWrappingError
		if errors.As(err, &mwe) {
			fmt.Println("Error as MyWrappingError")
		} else {
			fmt.Println("Error as NOT MyWrappingError")
		}
		fmt.Println()

		// Exit with value 1, since an error occurred
		fmt.Println("Exiting... ")
		os.Exit(1)
	}
	fmt.Println("No error 🙂")
	fmt.Println("Result: ", result)
	fmt.Println()
}

// Allow only numerators between -100 and 100, and only non-zero and positive even denominators smaller than 10
func increment(x int) (int, error) {
	// By convention, the last return value is of type error
	if x == 1 {
		// Using errors.New to create an error
		return 0, errors.New("x == 1")
	}
	if x == 2 {
		// Using fmt.Errorf to create an error
		return 0, fmt.Errorf("x == %d", x)
	}
	if x == 3 {
		// Using a sentinel error
		return 0, ErrSentinel
	}
	if x == 4 {
		// Using my custom error
		return 0, MyError("x == 4")
	}
	if x == 5 {
		// Using my other custom error, with an error code
		return 0, MyOtherError{99, "x == 5"}
	}
	if x == 6 {
		// Using fmt.Errorf with %w to wrap an error
		e := errors.New("wrapped error")
		return 0, fmt.Errorf("x == %d, root cause: %w", x, e)
	}
	if x == 7 {
		// Using a custom wrapping error
		e := MyError("wrapped error")
		return 0, MyWrappingError{"x == 7", e}
	}
	if x == 8 {
		// Using a sentinel wrapping error
		return 0, MyWrappingError{"x == 8", ErrSentinel}
	}
	if x == 9 {
		// Using two levels of wrapping
		e := fmt.Errorf("wrapped error wrapping error: %w", ErrSentinel)
		return 0, MyWrappingError{"x == 9", e}
	}
	if x == 10 {
		// Using multiple errors
		err1 := errors.New("First error")
		err2 := errors.New("Second error")
		err3 := errors.New("Third error")
		return 0, errors.Join(err1, err2, err3)
	}
	return x + 1, nil
}

// Sentinel errors are declared at package level
// Convention is for sentinel error names to start with "Err"
// Note that error messages must not be capitalized nor end with punctuation or a newline
var ErrSentinel = errors.New("sentinel error")

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
