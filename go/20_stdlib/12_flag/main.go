package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Go Standard Library: flag")
	fmt.Println()

	var name string
	var age int
	var vaccinated bool

	// Use functions that return pointers...
	//name := flag.String("name", "Alice", "Your name")
	//age := flag.Int("age", 42, "Your age")
	//vaccinated := flag.Bool("bptr", true, "Are you vaccinated?")

	// Or use functions that take pointers
	flag.StringVar(&name, "name", "Alice", "Your name")
	flag.IntVar(&age, "age", 21, "Your age")
	flag.BoolVar(&vaccinated, "vaccinated", true, "Are you vaccinated?")

	// Once all flags are declared, parse them...
	flag.Parse()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Vaccinated? %t\n", vaccinated)
	fmt.Println()

	fmt.Println("tail:", flag.Args())
	fmt.Println()
}
