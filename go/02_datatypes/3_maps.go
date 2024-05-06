package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")
	fmt.Println()

	// Maps
	// A map is similar to a slice, except it can be indexed by any comparable type, and there is no order.
	// Since the key must be a comparable type, you cannot use a slice or map as the key.
	var empty map[string]int // Default nil value for a slice is nil (same as slices)
	fmt.Printf("empty=%v (%T)\n", empty, empty)
	// Like slices, maps cannot be compared to each other, but they can be compared to nil
	fmt.Printf("empty == nil? %t\n", empty == nil)
	// Creating a map using an empty map literal results in an empty but non-nil map
	var nonNil = map[string]int{}
	fmt.Printf("nonNil=%v (%T)\n", nonNil, nonNil)
	fmt.Printf("nonNil == nil? %t\n", nonNil == nil)
	fmt.Println()

	// Creating a map using a non-empty map literal
	var students = map[string]int{
		"Alice": 21,
		"Bob":   65,
		"Carol": 42, // Note that each entry has a comma - even the last one (making it easy to add/remove entries)
	}
	fmt.Printf("students=%v (%T)\n", students, students)
	// Creating a map using make, and specifying an initial size
	pets := make(map[string]bool, 10)
	fmt.Println()

	// The built-in "len" function returns the number of key-value pairs
	fmt.Printf("len(empty)=%d\n", len(empty))
	fmt.Printf("len(students)=%d\n", len(students))
	// While pets has an initial size, len returns the number of actual entries
	fmt.Printf("len(pets)=%d\n", len(pets))
	// The built-in "cap" function doesn't work for maps...
	//fmt.Printf("cap(pets)=%d\n", cap(pets))
	fmt.Println()

	// Setting a map entry is similar to setting an array or slice entry
	pets["Bruno"] = false
	pets["Felix"] = false
	pets["Fido"] = true
	pets["Felix"] = true // Overrides the previous value
	fmt.Printf("pets=%v (%T)\n", pets, pets)
	fmt.Printf("len(pets)=%d\n", len(pets)) // Now pets has a len > 0 but < initial size
	fmt.Println()

	// Getting a map entry is similar to getting an array or slice entry
	fmt.Printf("pets[\"Bruno\"]=%v (%T)\n", pets["Bruno"], pets["Bruno"])
	fmt.Printf("pets[\"Felix\"]=%v (%T)\n", pets["Felix"], pets["Felix"])
	fmt.Printf("pets[\"Fido\"]=%v (%T)\n", pets["Fido"], pets["Fido"])
	// Getting a map entry that doesn't exist returns the nil value for the key type
	fmt.Printf("pets[\"Trixie\"]=%v (%T)\n", pets["Trixie"], pets["Trixie"])
	fmt.Println()

	// But if we get the nil value for a key (in this case, false),
	// is it because it doesn't exist or because it is actually false?
	// The comma-ok idiom answers this question
	// Since we know "Bruno" and "Trixie" both return false...
	_, ok := pets["Bruno"]
	if ok {
		fmt.Println("Bruno exists")
	} else {
		fmt.Println("Bruno does not exist")
	}
	_, ok = pets["Trixie"]
	if ok {
		fmt.Println("Trixie exists")
	} else {
		fmt.Println("Trixie does not exist")
	}
	fmt.Println()

	// The built-in "delete" function removes map entries
	delete(pets, "Felix")
	fmt.Printf("pets=%v (%T)\n", pets, pets)
	// The delete function does not return anything, and deleting a non-existent entry has no effect
	delete(pets, "Trixie")
	fmt.Println()

	// To iterate through map entries...
	for k, v := range students {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
	for k, v := range pets {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}
