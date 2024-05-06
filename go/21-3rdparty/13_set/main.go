package main

import (
	"fmt"

	"github.com/golang-collections/collections/set"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	fmt.Println("Go 3rd-party Library: collections")
	fmt.Println()

	fmt.Println("============ Custom set using map of strings to empty structs ============")
	var empty struct{}
	fmt.Println("Creating a new set...")
	var customset = map[string]struct{}{
		"string-1": empty,
		"string-2": empty,
	}
	fmt.Printf("customset: %v\n", customset)
	fmt.Printf("length: %v\n", len(customset))

	fmt.Println("Adding strings...")
	customset["string-2"] = empty // Adding the same element more than once has no effect
	customset["string-3"] = empty
	fmt.Printf("customset: %v\n", customset)
	fmt.Printf("length: %v\n", len(customset))

	fmt.Println("Removing string...")
	delete(customset, "string-2")
	fmt.Printf("customset: %v\n", customset)
	fmt.Printf("length: %v\n", len(customset))

	fmt.Println("Checking if string exists...")
	// Use the comma-ok idiom...
	_, ok := customset["string-1"]
	fmt.Printf("exists? %t\n", ok)
	_, ok = customset["string-2"]
	fmt.Printf("exists? %t\n", ok)
	_, ok = customset["string-3"]
	fmt.Printf("exists? %t\n", ok)

	fmt.Println()

	fmt.Println("============ github.com/golang-collections/collections/set ============")
	fmt.Println("Creating a new set...")
	set := set.New("string-1", "string-2")
	fmt.Printf("set: %v\n", set)
	fmt.Printf("length: %v\n", set.Len())

	fmt.Println("Adding strings...")
	set.Insert("string-2") // Adding the same element more than once has no effect
	set.Insert("string-3")
	fmt.Printf("set: %v\n", set)
	fmt.Printf("length: %v\n", set.Len())

	fmt.Println("Removing string...")
	set.Remove("string-2")
	fmt.Printf("set: %v\n", set)
	fmt.Printf("length: %v\n", set.Len())

	fmt.Println("Checking if string exists...")
	fmt.Printf("exists? %t\n", set.Has("string-1"))
	fmt.Printf("exists? %t\n", set.Has("string-2"))
	fmt.Printf("exists? %t\n", set.Has("string-3"))
	fmt.Println()

	fmt.Println("============ github.com/deckarep/golang-set/v2 ============")
	fmt.Println("Creating a new mapset...")
	mapset := mapset.NewSet[string]("string-1", "string-2")
	fmt.Printf("mapset: %v\n", mapset)
	// No length/size method, so convert to slice
	fmt.Printf("length: %v\n", len(mapset.ToSlice()))

	fmt.Println("Adding strings...")
	mapset.Add("string-2") // Adding the same element more than once has no effect
	mapset.Add("string-3")
	fmt.Printf("mapset: %v\n", mapset)
	fmt.Printf("length: %v\n", len(mapset.ToSlice()))

	fmt.Println("Removing string...")
	mapset.Remove("string-2")
	fmt.Printf("mapset: %v\n", mapset)
	fmt.Printf("length: %v\n", len(mapset.ToSlice()))

	fmt.Println("Checking if string exists...")
	fmt.Printf("exists? %t\n", mapset.Contains("string-1"))
	fmt.Printf("exists? %t\n", mapset.Contains("string-2"))
	fmt.Printf("exists? %t\n", mapset.Contains("string-3"))
	fmt.Println()

	fmt.Println("Done.")

}
