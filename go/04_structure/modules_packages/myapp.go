package main

import (
	"fmt"
	"github.com/hackebrot/turtle"
	"mymodule/mypackage1"
	"mymodule/mypackage2"
)

func main() {
	fmt.Println("main...")
	fmt.Println()

	// Local packages
	fmt.Println(mypackage1.Foo())
	fmt.Println(mypackage1.Bar())
	// Function "bar" is not visible outside of its package
	//fmt.Println(mypackage1.bar())
	fmt.Println(mypackage1.Foobar())
	fmt.Println()

	fmt.Println(mypackage2.Fox(), mypackage2.Dog())
	fmt.Println()

	// 3rd party dependency
	emojis := turtle.Search("computer")
	if emojis == nil {
		fmt.Println("No emojis found!")
	} else {
		fmt.Printf("Found emojis: %s\n", emojis)
	}
}
