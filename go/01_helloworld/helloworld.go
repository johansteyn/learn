// Naming the package "main" makes it an executable (as opposed to a library).
package main

// The go fmt tool will sort imports alphabetically
import (
	"fmt"
	"os"
)

// Execution begins in the "main" function
func main() {
	//fmt.Println("os.Args: ", os.Args)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Hello World!")
	} else {
		fmt.Println("Hello " + args[0] + "!")
	}
}
