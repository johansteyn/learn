package main

import (
	"fmt"
)

// Format of function declaration:
//  func <name>([<params...>]) <return types> {...}

// The entry function for Go programs is called "main"
func main() {
	fmt.Println("Functions")
	fmt.Println()

	foo()
	fmt.Printf("bar(7) returned %d\n", bar(7))
	x, y := 7, 12
	x, y = foobar(x, y)
	fmt.Printf("foobar(7, 12) returned %d and %d\n", x, y)
	x, y = barfoo(x, y)
	fmt.Printf("barfoo(12, 7) returned %d and %d\n", x, y)
	x, y = barfly(x, y)
	fmt.Printf("barfly returned %d and %d\n", x, y)
	fmt.Println()

	// Variadic parameters
	// Can pass any number of parameters, including zero
	print()
	print(1)
	print(1, 2, 7, 12, 24)
	// Can also pass a slice - note the dots...
	odds := []int{1, 3, 5, 7, 9}
	print(odds...)
	// Non-variadic parameters must be passed before variadic ones
	// In this case the first parameter (10) is non-variadic.
	inxs := inc(10, 1, 2, 7, 12, 24)
	fmt.Printf("inxs: %v\n", inxs)
	fmt.Println()

	// Functions are values
	var functions = []func(int, int) int{add, sub, mul, div}
	for i, f := range functions {
		fmt.Printf("function #%d: %d\n", i, f(9, 3))
	}
	fmt.Println()

	// Same as above, but using a "type" for readability
	type mathfunc func(int, int) int
	mathfuncs := []mathfunc{add, sub, mul, div}
	for i, f := range mathfuncs {
		fmt.Printf("function #%d: %d\n", i, f(9, 3))
	}
	fmt.Println()

	// Anonymous function
	n := func(a int, b int) int {
		return a / b
	}(42, 7) // Invoked immediately with parameter values 42 and 7
	fmt.Printf("Anonymous function returned %d\n", n)
	// This in itself isn't useful - anonymous functions are more useful with "defer" (TODO...)
	fmt.Println()

	// Closures
	f := outer(7)
	fmt.Printf("Inner function returns %d\n", f())
	fmt.Printf("Inner function returns %d\n", f())
	fmt.Printf("Inner function returns %d\n", f())
	fmt.Println()

	// Defer
	dirty()
	fmt.Println()

	// Recursion
	recursive(9)
}

// Takes no parameters and doesn't return anything
func foo() {
	fmt.Println("foo")
	// No return statement needed, unless returning before the end
}

// Takes one parameter and returns one value
func bar(a int) int {
	return a
}

// Takes two parameters and returns two values
// Note that multiple return values need to be in parentheses
// Also note that the return values are distinct - not "tuples" as in Python or Scala
func foobar(a, b int) (int, int) {
	x := b
	y := a
	return x, y
}

// Same as above, but return variables can be named
func barfoo(a, b int) (x, y int) {
	x = b
	y = a
	return x, y
}

// Same as above, using a blank return
// Not recommended as code is less readable
func flybar(a, b int) (x, y int) {
	x = b
	y = a
	// No need to specify the return values here
	return
}

// Same as above, but without actually using the return variables at all
// Returned values are automatically assigned to the return variables
// So, named return values are not very useful, but they are needed with "defer" (TODO...)
// Again, not recommended - pointless naming return variables that are not used
// Q: Why does Go even allow this?
//
//	Unused variables normally result in compile errors...
func barfly(a, b int) (x, y int) {
	return b, a
}

// Takes a variable number of parameters (variadic parameters)
// A variadic parameter is essentially a slice, constructed from the parameters
func print(xs ...int) {
	fmt.Printf("Printing %d parameters...\n", len(xs))
	for i, x := range xs {
		fmt.Printf("#%d: %d\n", i, x)
	}
}

// At most one parameter can be variadic, and it must be the last one
func inc(value int, xs ...int) []int {
	inxs := make([]int, 0, len(xs))
	for _, x := range xs {
		inxs = append(inxs, x+value)
	}
	return inxs
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// If both parameters are of the same type we only need the type once
func mul(a, b int) int {
	return a * b
}

// Note that the return type can be in parentheses
func div(a int, b int) int {
	return a / b
}

// Function that returns an anonymous inner function
func outer(x int) func() int {
	return func() int {
		// The inner function "captures" parameter x, forming a "closure"
		x += 1
		return x
	}
}

func cleanup1() {
	fmt.Println("Cleanup1")
}

func cleanup2() {
	fmt.Println("Cleanup2")
}

// Deferred functions are only called when this function returns
// They are called in reverse order (from last to first),
// as per page 144 of "The Go Programming Language"
func dirty() {
	// Defering directly
	defer cleanup1()
	defer cleanup2()

	// Defering via local variables
	f := cleanup1
	defer f()
	f = cleanup2
	defer f()

	// Defering via anonymous functions (so that we can do some extra work)
	f = cleanup1
	defer func() {
		fmt.Println("Extra1")
		// Careful... it will call cleanup2 - not cleanup1
		f()
	}()
	f = cleanup2
	defer func() {
		fmt.Println("Extra2")
		f()
	}()

	defer func(f func()) {
		fmt.Println("Extra1")
		f()
	}(cleanup1)
	defer func(f func()) {
		fmt.Println("Extra2")
		f()
	}(cleanup2)

	fmt.Println("Dirty")
}

func recursive(i int) {
	if i > 0 {
		fmt.Printf("%d ", i)
		recursive(i - 1)
	} else {
		fmt.Println()
	}
}
