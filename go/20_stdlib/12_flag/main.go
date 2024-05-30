package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Go Standard Library: flag")
	fmt.Println()

	flags1()
	//flags2()

	fmt.Println("tail:", flag.Args())
	fmt.Println()
}

// Call functions that return pointers
func flags1() {
	namePtr := flag.String("name", "Alice", "Your name")
	agePtr := flag.Int("age", 21, "Your age")
	vaccinatedPtr := flag.Bool("vaccinated", true, "Are you vaccinated?")
	// Once all flags are declared, parse them...
	flag.Parse()
	// Only assign the values after parsing
	name := *namePtr
	age := *agePtr
	vaccinated := *vaccinatedPtr

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Vaccinated? %t\n", vaccinated)
}

// Call functions that take pointers
func flags2() {
	var name string
	var age int
	var vaccinated bool

	flag.StringVar(&name, "name", "Alice", "Your name")
	flag.IntVar(&age, "age", 21, "Your age")
	flag.BoolVar(&vaccinated, "vaccinated", true, "Are you vaccinated?")
	// Once all flags are declared, parse them...
	flag.Parse()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Vaccinated? %t\n", vaccinated)
	fmt.Println()
}
