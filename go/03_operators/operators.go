package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operators (TODO...")
	fmt.Println()

  // https://www.youtube.com/watch?v=0c-1KJwSMCw
	// At 3:15 he says the result is 8, but I get 4...
	fmt.Println(1 << 2)
	i := 2
	fmt.Println(1 << i)
	i64 := int64(2)
	fmt.Println(1 << i64)

	// You can only shift a non-negative amount of places.
	// In previous versions of Go a negative value would be cast to an unsigned int,
	// which would work but for an unexpected result...
	// Since Go 1.13 specifying a negative number will simply result in a panic
	// (which is better that silently not working as expected...)
	//fmt.Println(1 << -2)

}

