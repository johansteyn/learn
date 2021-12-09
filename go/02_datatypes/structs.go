package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")

	// Declare a struct with the "type" keyword, the struct name, and the "struct" keyword,
	// followed by braces containing the fields
	type person struct {
		name string // Note that there are no commas between struct fields
		age int
		vaccinated bool
	}
	// To define a variable of that type
	var alice person // Default nil value, where each field has its nil value
	fmt.Printf("alice=%v (%T)\n", alice, alice)
	// TIP: Using %+v shows more info
	fmt.Printf("alice=%+v (%T)\n", alice, alice)
	// Using an empty struct literal
	bob := person{}
	fmt.Printf("bob=%v (%T)\n", bob, bob)
	// Unlike slices and maps, there is no such thing as a nil struct
	// ie. alice and bob are both empty - not nil
	// So, we cannot compare alice or bob to nil 
	//fmt.Printf("alice == nil? %t\n", alice == nil)
	//fmt.Printf("bob == nil? %t\n", bob == nil)
	// But we can compare alice and bob to each other
	fmt.Printf("alice == bob? %t\n", alice == bob)
	// Another way to create an empty variable is with the "new" keyword
	// This actually creates a pointer, so we need to dereference it
	carol := new(person)
	fmt.Printf("*carol=%v (%T)\n", *carol, *carol)
	fmt.Printf("alice == *carol? %t\n", alice == *carol)

	// Struct literals can list all field values, in order
	dave := person{"Dave", 36, true}
	fmt.Printf("dave=%v (%T)\n", dave, dave)
	// Or any individual fields can be named in any order
	edgar := person{
		age: 18, // Note that here we do have commas
		name: "Edgar", // The last field also has a comma
	}
	fmt.Printf("edgar=%v (%T)\n", edgar, edgar)
	// Individual fields are accessed (set/get) using dot notation
	alice.name = "Alice"
	alice.age = 21
	fmt.Printf("alice=%v (%T)\n", alice, alice)
	bob.name = "Bob"
	bob.age = 65
	fmt.Printf("bob=%v (%T)\n", bob, bob)
	// Note that we use dot notation with pointers too (there is no -> operator in Go)
	carol.name = "Carol"
	carol.age = 39
	fmt.Printf("*carol=%v (%T)\n", *carol, *carol)

	// An anonymous struct does not use the "type" keyword
	var bruno struct {
		name string
		age int
		vaccinated bool
	}
	bruno.name = "Bruno"
	fmt.Printf("bruno=%v (%T)\n", bruno, bruno)
	// In this case our anonymous struct has the exact same fields as the person struct,
	// so we can compare it to any person.
	fmt.Printf("alice == bruno? %t\n", alice == bruno)
}
