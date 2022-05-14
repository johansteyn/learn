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

	// However, that's not what we really want...
	// We want to call a proper conversion function that calculates the correct value
	var calculatedF = celsiusToFahrenheit(freezingC)
	fmt.Printf("calculatedF = %d\n", calculatedF)
	if freezingF == calculatedF {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", freezingF, freezingC)
	} else {
		fmt.Printf("%d Celsius != %d Fahrenheit\n", freezingF, freezingC)
	}
	fmt.Println()
}

type celsius int
type fahrenheit int

func celsiusToFahrenheit(c celsius) fahrenheit {
	var x float32 = float32(c) * 9.0 / 5.0 + 32.0
	return fahrenheit(x)
}

func fahrenheitToCelsius(f fahrenheit) celsius  {
	var x float32 = 5.0 / 9.0 * (float32(f) - 32.0)
	return celsius(x)
}

