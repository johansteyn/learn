package main

import (
	// Everything imported is in the "file" block
	// ie. they have scope only in this file
	"fmt"
)

// Anything declared outside of a fuction is in the "package" block
// ie. it has scope across all files in the same package
const freezingTemp = 0
const boilingTemp = 100

func main() {
	fmt.Println("Blocks")

	// Here we print the value of the package scope constant
	fmt.Println("boilingTemp:", boilingTemp)
	// And here we "shadow" it...
	var boilingTemp = 273
	fmt.Println("boilingTemp:", boilingTemp)
	if boilingTemp > 100 {
		// Here we "shadow" both!
		boilingTemp := -42
		fmt.Println("boilingTemp:", boilingTemp)
	}
	// Back to the local variable
	fmt.Println("boilingTemp:", boilingTemp)

	// There is also a "universe" block where some identifiers are pre-declared, eg:
	//   Types: int, string, etc.
	//   Constants: true, false, etc.
	//   Functions: make, close, etc.
	//   Stuff like nil
	// So, it's possible to shadow them...
	fmt.Println("true:", true)
	var true = 42
	fmt.Println("true:", true)
	var int = 1
	fmt.Println("int:", int)
	make := "Shadowing the built-in 'make' function..."
	fmt.Println("make:", make)

	// Shadowing is not detected by the compiler, vet or lint.
	// The "shadow" linter is able to detect shadowing.
	// To install:
	//   % go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	// To run:
	//   % shadow <file.go>
}
