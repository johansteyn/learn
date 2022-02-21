package main

import (
	"fmt"
	"math"
)

// Format of method declaration: 
//  func (<id> <receiver>) <name>([<params...>]) <return types> {...}
// ie. essentially a function with a receiver

func main() {
	fmt.Println("Methods")
	fmt.Println()

	rect := rectangle{
		width: 7,
		height: 9,
	}
	fmt.Printf("rect.area(): %d\n", rect.area())

	circ := circle{
		radius: 5,
	}
	// The method has a value receiver, so use a value
	fmt.Printf("circ.area(): %f\n", circ.area())
	// But we can also invoke the method on a pointer.
	// The compiler will implicitly dereference the pointer.
	circ_ptr := &circ
	fmt.Printf("circ_ptr.area(): %f\n", circ_ptr.area()) // implicit *circ_ptr

	ctr := counter{
		total: 12,
	}
	fmt.Printf("ctr.total: %d\n", ctr.total)
	// The method has a pointer receiver, so use a pointer
	ctr_ptr := &ctr
	ctr_ptr.inc(2)
	fmt.Printf("ctr.total: %d\n", ctr.total)
	// But we can also invoke the method directly on the value.
	// The compiler will implicitly take the address of the variable.
	ctr.inc(3)
	fmt.Printf("ctr.total: %d\n", ctr.total)

	// Caution: invoking the method on a copy of the value has no effect
	// Eg: passing a parameter by value makes a copy of the variable
	incorrectIncFive(ctr)
	fmt.Printf("ctr.total: %d\n", ctr.total)
	// Instead, pass a pointer in order to reference the original
	correctIncFive(&ctr)
	fmt.Printf("ctr.total: %d\n", ctr.total)

	// We can invoke a method on a nil pointer
	correctIncFive(nil)

	// We can also invoke the method on an unassigned variable
	// because it will have the default nil value for a struct
	// which is a struct whose fields all have the default nil value.
	// In this case the "total" field has default nil value of zero
	var c counter
	fmt.Printf("c.total: %d\n", c.total)
	c.inc(7)
	fmt.Printf("c.total: %d\n", c.total)

	// Method values and method expressions
	// A method value is similar to a closure in that it gets
	// bound to an instance, whose value persists between calls
	foo := rect.area // Method value with type: func()int
	fmt.Printf("foo(): %d\n", foo()) // Equivalent to: rect.area()
	// A method expression is obtained from the type
	bar := rectangle.area // Method expression with type: func(rectangle)int
	// The first parameter must be the receiver for the method
	// In that way, a method becomes just like a function
	fmt.Printf("bar(rect): %d\n", bar(rect)) // Also equivalent to: rect.area()
}

//----------------------------------------
type rectangle struct {
	width, height int
}

// Methods must be declared in the same package as the receiver type
// They can be in a different file, but it's best to have them in the same file
// Idiomatic Go uses short receiver names, and never "this" or "self"
func (r rectangle) area() int {
	return r.width * r.height
}

//----------------------------------------
type circle struct {
	radius float32
}

// Go does not have method overloading
// But you can use the same method name for different types
func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

//----------------------------------------
type counter struct {
	total int
}

// The "area" methods above both use value receivers
// To modify the receiver, use a pointer receiver
// As with parameters, a pointer receiver enables us to modify the original
func (c *counter) inc(n int) {
	// Also, like pointer parameters, a pointer receiver can be nil
	// It is then up to us to do the nil check
	if c == nil {
		fmt.Println("Nil pointer receiver!")
		return
	}
	// Note that we don't need to dereference the receiver - even though it's a pointer
	c.total += n
}

// Using a value parameter we invoke the method on a copy, which has no effect
func incorrectIncFive(c counter) {
	c.inc(5)
}

// Using a pointer parameter, we invoke the method on the original
func correctIncFive(c *counter) {
	c.inc(5)
}

