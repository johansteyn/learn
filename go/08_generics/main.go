package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Printf("Generics\n")
	fmt.Println()

	ints := []int{1, 2, 3}
	fmt.Printf("sumInts: %d\n", sumInts(ints))
	fmt.Println()

	floats := []float32{1.0, 2.0, 3.0}
	fmt.Printf("sumFloats: %f\n", sumFloats(floats))
	fmt.Println()

	fmt.Printf("sumIntsOrFloats[int]: %d\n", sumIntsOrFloats(ints))
	fmt.Println()

	fmt.Printf("sumIntsOrFloats[float32]: %f\n", sumIntsOrFloats(floats))
	fmt.Println()

	var ageAlice age = 21
	var ageBob age = 64
	var ageCarol age = 42
	ages := []age{ageAlice, ageBob, ageCarol}
	fmt.Printf("sumIntsOrFloats2[age]: %d\n", sumIntsOrFloats2(ages))
	fmt.Println()


	fmt.Printf("sumNumbers[int]: %d\n", sumNumbers(ints))
	fmt.Println()

	fmt.Printf("sumNumbers[float32]: %f\n", sumNumbers(floats))
	fmt.Println()

	bytes := []byte{1, 2, 3}
	fmt.Printf("sumNumbers[bytes]: %d\n", sumNumbers(bytes))
	fmt.Println()

	fmt.Printf("sumNumbers2[int]: %d\n", sumNumbers2(ints))
	fmt.Println()

	fmt.Printf("sumNumbers2[float32]: %f\n", sumNumbers2(floats))
	fmt.Println()

	fmt.Printf("sumNumbers2[bytes]: %d\n", sumNumbers2(bytes))
	fmt.Println()

	fmt.Println("Done.")
}

func sumInts(numbers []int) int {
	var result int
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// sumFloats is almost identical to sumInts
// The only difference is the types...
func sumFloats(numbers []float32) float32 {
	var result float32
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// sumIntsOrFloats is generic - it can work with ints or floats
func sumIntsOrFloats[T int | float32](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

type age int

// sumIntsOrFloats2 can work with ints or floats or any type that boils down to either
func sumIntsOrFloats2[T ~int | ~float32](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// A Number can be of type int or float32 or even byte
type Number interface {
	int | float32 | byte
}

// sumNumbers has a "cleaner" signature
func sumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

// sumNumbers2 has an even cleaner signature
func sumNumbers2[T constraints.Ordered](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

