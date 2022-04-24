// Naming the package "main" makes it an executable (as opposed to a library).
package main

// The go fmt tool will sort imports alphabetically
import (
	"fmt"
	// In case of name conflicts, an import can be renamed
	newname "os"
	// Any unused import results in a compile error
	// To avoid that, use a blank import name
	_ "runtime"
)

// Execution begins in the "main" function
func main() {
	// Here we use "newname" instead of "os"
	args := newname.Args[1:]
	if len(args) == 0 {
		fmt.Println("Hello World!")
	} else {
		fmt.Println("Hello " + args[0] + "!")
	}
}
