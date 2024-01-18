package main

import (
	"fmt"
	"os"
)

// https://go.dev/doc/articles/race_detector
// Accidentally shared variable
func main() {
	data := [...]byte{0, 1, 2}
	ParallelWrite(data[:])
	res := <-ParallelWrite(data[:])
	fmt.Printf("*** Result: %+v\n", res)
}

// Writes data concurrently to file1 and file2 assigning the returned error to the same variable.
// The fix is to use := instead of = so that errors are assigned to separate variables
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		fmt.Printf("Error creating file1: %v\n", err)
		res <- err
	} else {
		go func() {
			_, err = f1.Write(data) // The first conflicting write to err.
			//_, err := f1.Write(data) // Fixed by using :=
			if err != nil {
				fmt.Printf("Error writing to file1: %v\n", err)
			}
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2")
	if err != nil {
		fmt.Printf("Error creating file2: %v\n", err)
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data) // The second conflicting write to err.
			//_, err := f2.Write(data) // Fixed by using :=
			if err != nil {
				fmt.Printf("Error writing to file2: %v\n", err)
			}
			res <- err
			f2.Close()
		}()
	}
	return res
}
