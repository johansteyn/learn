package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")
	fmt.Println()

	// Arrays
	fmt.Println("------------ Arrays ------------")
	// All elements must be of the same type
	// Must specify size, which cannot change (ie. size forms part of its type)
	var powersOfTwo [10]int // Array of 10 integers, all with default nil values (zero)
	fmt.Printf("powersOfTwo=%v (%T)\n", powersOfTwo, powersOfTwo)
	powersOfTwo[0] = 2
	for i := 1; i < len(powersOfTwo); i++ {
		powersOfTwo[i] = 2 * powersOfTwo[i-1]
	}
	fmt.Printf("powersOfTwo=%v (%T)\n", powersOfTwo, powersOfTwo)
	fmt.Println()

	// An array can be initialized using an array literal
	var names = [5]string{"Alice", "Bob", "Carol", "Dave", "Edgar"}
	fmt.Printf("names=%v (%T)\n", names, names)
	// Sparse array: leaving out the ladies' ages
	var ages = [5]int{1: 19, 3: 28, 65}
	fmt.Printf("ages=%v (%T)\n", ages, ages)
	// Size is not needed when assigning an array literal
	var vegetarians = [...]bool{2: true, false, true}
	fmt.Printf("vegetarians=%v (%T)\n", vegetarians, vegetarians)
	// But if a size is specified it cannot be less than the literal size
	//var idiots = [3]bool{false, true, true, false, true}
	// It may be larger though - unassigned elements will have the nil value
	var idiots = [6]bool{false, true, true, false, true}
	fmt.Printf("idiots=%v (%T)\n", idiots, idiots)
	fmt.Println()

	// The builtin "len" function returns an array's size
	fmt.Printf("Size of names array=%d\n", len(names))
	// An array's type includes its size, which means that:
	// - You cannot assign an array of type [3]int to an array of type [5]int
	// - You cannot use a varaiable to declare the array size, as it must be known at compile time
	// Therefore, arrays are seldom used - use arrays only when you know the exact size needed up front.
	// Arrays are mainly used as the backing store for slices.
	fmt.Println()

	// Slices
	fmt.Println("------------ Slices ------------")
	// A slice is similar to an array, but the size does not form part of its type
	// Using [n] or [...] declares an array, while using [] declares a slice
	// Since slices can grow and shrink, you can have an empty slice (but not an empty array).
	var nilSlice []int // Default nil value for a slice is... nil
	fmt.Printf("nilSlice=%v (%T)\n", nilSlice, nilSlice)
	// Go's nil is not quite like C or Java's null
	// It's an identifier that represents the lack of a value.
	// Like literals and untyped constants, nil has no type so can be assigned or compared to values of different types.
	// Slices aren't comparable, ie. you can't use == or !=, but you can compare a slice to nil.
	fmt.Printf("nilSlice == nil? %t\n", nilSlice == nil)

	// With a nil slice you can get the length and capacity and iterate over it without getting a nil pointer exception
	fmt.Println("Working with a nil slice...")
	fmt.Printf("len(nilSlice)=%d, cap(nilSlice)=%d\n", len(nilSlice), cap(nilSlice))
	for i, x := range nilSlice {
		fmt.Printf("nilSlice[%d]=%d\n", i, x)
	}
	// But if you reference a nil slice you will get a runtime panic...
	//fmt.Printf("nilSlice[0]: %d\n", nilSlice[0])

	// Creating a slice using an empty slice literal results in an empty but non-nil slice
	var empty = []int{}
	fmt.Printf("empty=%v (%T)\n", empty, empty)
	fmt.Printf("empty == nil? %t\n", empty == nil)
	fmt.Println()

	// Grow slices with the built-in "append" function
	nonEmpty := append(empty, 1)
	fmt.Printf("nonEmpty=%v (%T)\n", nonEmpty, nonEmpty)
	// Appending three elements (can append any number)
	nonEmpty = append(nonEmpty, 2, 3, 4)
	fmt.Printf("nonEmpty=%v (%T)\n", nonEmpty, nonEmpty)
	var pets = []string{"Bruno", "Felix", "Fido"}
	fmt.Printf("pets=%v (%T)\n", pets, pets)
	pets = append(pets, "Koos", "Pluto", "Bernard", "Gizmo")
	fmt.Printf("pets=%v (%T)\n", pets, pets)
	// Like arrays, slices can also be sparse
	var cats = []bool{1: true, 3: true, false, 6: true}
	fmt.Printf("cats=%v (%T)\n", cats, cats)
	// An entire slice can be appended to a slice using the ... operator to expand the slice being added
	var semester1 = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}
	var semester2 = []string{"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Printf("semester1=%v (%T)\n", semester1, semester1)
	fmt.Printf("semester2=%v (%T)\n", semester2, semester2)
	var year = append(semester1, semester2...)
	fmt.Printf("year=%v (%T)\n", year, year)
	fmt.Println()

	// In addition to size (len), a slice also has capacity (cap)
	fmt.Printf("len(nilSlice)=%d, cap(nilSlice)=%d\n", len(nilSlice), cap(nilSlice))
	fmt.Printf("len(empty)=%d, cap(empty)=%d\n", len(empty), cap(empty))
	fmt.Printf("len(nonEmpty)=%d, cap(nonEmpty)=%d\n", len(nonEmpty), cap(nonEmpty))
	fmt.Printf("len(pets)=%d, cap(pets)=%d\n", len(pets), cap(pets))
	fmt.Printf("len(cats)=%d, cap(cats)=%d\n", len(cats), cap(cats))
	fmt.Printf("len(semester1)=%d, cap(semester1)=%d\n", len(semester1), cap(semester1))
	fmt.Printf("len(year)=%d, cap(year)=%d\n", len(year), cap(year))
	// Capacity is increased when needed, doubling each time
	var numbers []int
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len(numbers)=%d, cap(numbers)=%d, numbers=%v \n", len(numbers), cap(numbers), numbers)
	}
	fmt.Println()

	// To create a slice with an initial length that you happen to know, use the built-in "make" function
	daysOfYear := make([]int, 365)
	fmt.Printf("len(daysOfYear)=%d, cap(daysOfYear)=%d, daysOfYear=%v \n", len(daysOfYear), cap(daysOfYear), daysOfYear)
	for i := 0; i < 365; i++ {
		daysOfYear[i] = i + 1
	}
	// Note that we didn't use "append" - we assigned values directly and the capacity remains at 365
	// Remember that append ALWAYS increase the size of a slice!
	fmt.Printf("len(daysOfYear)=%d, cap(daysOfYear)=%d, daysOfYear=%v \n", len(daysOfYear), cap(daysOfYear), daysOfYear)
	// Alternatively, we can make a slice with initial length zero but with known capacity, and then append
	daysOfLeapYear := make([]int, 0, 366)
	fmt.Printf("len(daysOfLeapYear)=%d, cap(daysOfLeapYear)=%d, daysOfLeapYear=%v \n", len(daysOfLeapYear), cap(daysOfLeapYear), daysOfLeapYear)
	daysOfLeapYear = append(daysOfLeapYear, daysOfYear...)
	daysOfLeapYear = append(daysOfLeapYear, 366)
	fmt.Printf("len(daysOfLeapYear)=%d, cap(daysOfLeapYear)=%d, daysOfLeapYear=%v \n", len(daysOfLeapYear), cap(daysOfLeapYear), daysOfLeapYear)
	// Now if we append just one more, the capacity will (almost) double...
	daysOfLeapYear = append(daysOfLeapYear, 367)
	fmt.Printf("len(daysOfLeapYear)=%d, cap(daysOfLeapYear)=%d, daysOfLeapYear=%v \n", len(daysOfLeapYear), cap(daysOfLeapYear), daysOfLeapYear)
	// Cannot specify a capacity that is smaller than the length
	//x := make([]int, 12, 7)
	fmt.Println()

	// A slice expression creates a slice from a slice (similar to Python)
	daysOfJanuary := daysOfYear[:31]    // 31 days, from index 0 to 30
	daysOfFebruary := daysOfYear[31:59] // 28 days, from index 31 to 58
	fmt.Printf("len(daysOfJanuary)=%d, cap(daysOfJanuary)=%d, daysOfJanuary=%v \n", len(daysOfJanuary), cap(daysOfJanuary), daysOfJanuary)
	fmt.Printf("len(daysOfFebruary)=%d, cap(daysOfFebruary)=%d, daysOfFebruary=%v \n", len(daysOfFebruary), cap(daysOfFebruary), daysOfFebruary)
	// Note how the capacities are from the start of the slice to the end of the original slice's capacity
	// Also note that if we change a value in the original slice, then it changes the value in the slice
	daysOfYear[0] = 999
	fmt.Printf("len(daysOfJanuary)=%d, cap(daysOfJanuary)=%d, daysOfJanuary=%v \n", len(daysOfJanuary), cap(daysOfJanuary), daysOfJanuary)
	// You can also take a slice of an array - handy if you need to pass a slice and you have an array.
	fiveToNine := numbers[4:9]
	fmt.Printf("len(fiveToNine)=%d, cap(fiveToNine)=%d, fiveToNine=%v \n", len(fiveToNine), cap(fiveToNine), fiveToNine)
	// But beware that it shares the underlying data with that array...
	fiveToNine[0] = 999
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Println()

	// The built-in "copy" function creates a slice that's independent of the original
	sixToEight := make([]int, 3)
	numCopied := copy(sixToEight, numbers[5:8]) // Could have used slice [5:] too - it will only copy 3 elements
	fmt.Printf("numCopied=%d\n", numCopied)
	fmt.Printf("len(sixToEight)=%d, cap(sixToEight)=%d, sixToEight=%v \n", len(sixToEight), cap(sixToEight), sixToEight)
	// Now, modifying the slice won't affect the original array
	sixToEight[1] = 666
	fmt.Printf("sixToEight=%v\n", sixToEight)
	fmt.Printf("numbers=%v\n", numbers)
	// The source and destination can be the same array/slice
	copy(numbers[4:], numbers[:4]) // Don't need to assign return value
	fmt.Printf("numbers=%v\n", numbers)
	// They can even overlap
	copy(numbers[2:], numbers[:4])
	fmt.Printf("numbers=%v\n", numbers)

	// To iterate through a slice, use a for loop with "range"
	fmt.Println("Iterating through a slice:")
	for i, number := range numbers {
		fmt.Printf(" #%d:%d\n", i, number)
	}
}
