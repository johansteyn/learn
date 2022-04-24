package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	fmt.Println()

	// Declare a struct type with:
	// - Keyword "type"
	// - The name of the new type
	// - Keyword "struct"
	// - Followed by braces containing the fields
	type person struct {
		name string // Note that there are no commas between struct fields
		age int
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
	fmt.Println()

	// Struct literals can list all field values, in order
	dave := person{"Dave", 36}
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
	fmt.Println()

	// An anonymous struct does not use the "type" keyword
	var frank struct {
		name string
		age int
	}
	frank.name = "Frank"
	fmt.Printf("frank=%v (%T)\n", frank, frank)
	// In this case our anonymous struct has the exact same fields as the person struct,
	// so we can compare it to any person.
	fmt.Printf("alice == frank? %t\n", alice == frank)
	fmt.Println()

	// Types can be composed by struct embedding
	// In OO design employee might extend person (is-a relationship)
	// and it might contain a status (has-a relationship)
	type status struct {
		vaccinated bool
	}
	type employee struct {
		title string
		person // Note that person and status are embedded fields
		status // Embedded fields don't need names - only types
	}
	george := person{"George", 93}
	vacc := status{true}
	storeman := employee{"Storeman", george, vacc}
	fmt.Printf("storeman=%v (%T)\n", storeman, storeman)
	// So, while Go does not have inheritance, it has a handy shortcut...
	// The fields of embedded structs can be referenced directly.
	// Referencing embedded fields indirectly
	fmt.Printf("storeman.person.name=%s\n", storeman.person.name)
	fmt.Printf("storeman.person.age=%d\n", storeman.person.age)
	fmt.Printf("storeman.status.vaccinated=%t\n", storeman.status.vaccinated)
	// Referencing embedded fields directly
	fmt.Printf("storeman.name=%s\n", storeman.name)
	fmt.Printf("storeman.age=%d\n", storeman.age)
	fmt.Printf("storeman.vaccinated=%t\n", storeman.vaccinated)
	fmt.Println()

	// But if two or more embedded structs have the same field...
	type pet struct {
		name string
		age int
	}
	type student struct {
		person
		pet
	}
	harry := person{"Harry", 13}
	benji := pet{"Benji", 7}
	headboy := student{harry, benji}
	// Direct references no longer work due to ambiguity
	//fmt.Printf("headboy.name=%s\n", headboy.name)
	//fmt.Printf("headboy.age=%d\n", headboy.age)
	// Indirect references still work fine
	fmt.Printf("headboy.person.name=%s\n", headboy.person.name)
	fmt.Printf("headboy.person.age=%d\n", headboy.person.age)
	fmt.Printf("headboy.pet.name=%s\n", headboy.pet.name)
	fmt.Printf("headboy.pet.age=%d\n", headboy.pet.age)

}
