package main

import (
	"fmt"
)

func main() {
	fmt.Println("Types")
	fmt.Println()

	// The type of a variable or expression defines:
	// - Its characteristics (eg: size, # elements, internal representation)
	// - Operations that can be performed on it (or methods associated with it)

	// Types are usually declared at package level (see below)
	// A type declaration defines a new named type for an existing underlying type
	// Form: type <name> <underlying-type>

	var freezingC celsius = 0
	var freezingF fahrenheit = 32

	// Cannot compare values of different types, even if their underlying types are the same
	//fmt.Println("freezingC==freezing? %t\n", freezingC == freezingF)

	// But can do a type conversion if their underlying types are the same
	var convertedF = fahrenheit(freezingC)
	fmt.Printf("convertedF = %d\n", convertedF)
	if freezingF == convertedF {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", freezingF, freezingC)
	} else {
		fmt.Printf("%d Fahrenheit != %d Celsius\n", freezingF, freezingC)
	}
	fmt.Println()

	var freezingCptr *celsius = &freezingC
	var freezingFptr *fahrenheit = &freezingF

	// Similarly, we cannot compare the contents of pointers to different types
	//fmt.Println("*freezingCptr==freezingFptr? %t\n", *freezingCptr == *freezingFptr)

	// And, again, we can do type conversion
	var convertedFptr *fahrenheit = (*fahrenheit)(freezingCptr)
	if *freezingFptr == *convertedFptr {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", *freezingFptr, *freezingCptr)
	} else {
		fmt.Printf("%d Fahrenheit != %d Celsius\n", *freezingFptr, *freezingCptr)
	}
	fmt.Println()

	// However, in this case type conversion is not what we really want...
	// We want to call a proper conversion function that calculates the correct value
	var calculatedF = celsiusToFahrenheit(freezingC)
	fmt.Printf("calculatedF = %d\n", calculatedF)
	if freezingF == calculatedF {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", freezingF, freezingC)
	} else {
		fmt.Printf("%d Celsius != %d Fahrenheit\n", freezingF, freezingC)
	}
	fmt.Println()

	/*
		var c int = 100
		var f int = 212
		var cptr *int = &c
		var fptr *int = &f
		var boilingCPtr *celsius = (*celsius)(cptr)
		var boilingFPtr *fahrenheit = (*fahrenheit)(fptr)
		if *boilingCPtr == *boilingFPtr {
			fmt.Printf("%d Celsius == %d Fahrenheit\n", *boilingCPtr, *boilingFPtr)
		} else {
			fmt.Printf("%d Celsius != %d Fahrenheit\n", *boilingCPtr, *boilingFPtr)
		}

		fmt.Printf("c = %d, *cptr = %d\n", c, *cptr)

		// Works...
		//var freezingCptr celsiusPtr = &c
		//var freezingCptr celsiusPtr = cptr
		//fmt.Printf("freezingCptr = %d\n", *freezingCptr)

		var freezingCptr *celsius = (*celsius)(cptr)
		fmt.Printf("freezingCptr = %d\n", *freezingCptr)
	*/
}

type celsius int
type fahrenheit int

func celsiusToFahrenheit(c celsius) fahrenheit {
	var x float32 = float32(c)*9.0/5.0 + 32.0
	return fahrenheit(x)
}

func fahrenheitToCelsius(f fahrenheit) celsius {
	var x float32 = 5.0 / 9.0 * (float32(f) - 32.0)
	return celsius(x)
}
