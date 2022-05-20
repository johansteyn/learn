package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	fmt.Println()

	// Simple pointer to int
	x := 7
	fmt.Printf("x=%d (%T)\n", x, x)
	xptr := &x // Address operator &
	fmt.Printf("xptr=%v (%T)\n", xptr, xptr)
	i := *xptr // Indirection operator * (contents of)
	fmt.Printf("*xptr=%v (%T)\n", *xptr, *xptr)
	fmt.Printf("i=%d (%T)\n", i, i)
	*xptr = 12 // Changing the contents of a pointer changes the original variable's value
	fmt.Printf("x=%d (%T)\n", x, x)
	fmt.Println()

	// Uninitialized pointer to string
	var sptr *string
	fmt.Printf("sptr=%v (%T)\n", sptr, sptr)
	// Is equal to nil
	fmt.Printf("sptr == nil? %v\n", sptr == nil)
	// And cannot be dereferenced until it is assigned a value
	//fmt.Printf("*sptr=%v (%T)\n", *sptr, *sptr)
	s := "The quick brown fox"
	sptr = &s
	fmt.Printf("sptr=%v (%T)\n", sptr, sptr)
	fmt.Printf("*sptr=%v (%T)\n", *sptr, *sptr)
	*sptr = "jumps over the lazy dog" // Again, changing the contents changes the original
	fmt.Printf("s=%s (%T)\n", s, s)
	fmt.Println()

	// Using "new" returns a pointer to a variable initialized to its zero value
	var nptr = new(int)
	fmt.Printf("nptr=%v (%T)\n", nptr, nptr)
	fmt.Printf("*nptr=%v (%T)\n", *nptr, *nptr)
	fmt.Println()

	// Can get a pointer to a struct literal using the address operator
	type Foo struct {
		x int
		s string
	}
	fooptr := &Foo{42, "The meaning of life"}
	// Why does it print the values and not the address?
	// Probably because it's the address of a struct literal...
	// ie. it was not created on the heap - it exists at compile time.
	fmt.Printf("fooptr=%v (%T)\n", fooptr, fooptr)
	foo := *fooptr
	fmt.Printf("foo=%v (%T)\n", foo, foo)
	fmt.Println()

	// But cannot take the address of a primitive literal or constant.
	const tobe = "To be or not to be"
	//sptr = &"Jumped over the lazy dog"
	//sptr = &tobe
	// Instead, assign the constant to a string variable and take its address
	s = tobe
	sptr = &s
	fmt.Printf("*sptr=%v (%T)\n", *sptr, *sptr)
	// Or use a function to obtain a pointer
	sptr = stringp(tobe)
	fmt.Printf("*sptr=%v (%T)\n", *sptr, *sptr)
	fmt.Println()

	// Go uses copy by value, 
	// ie. the the value passed to a function is copied to the parameter variable 
	// Therefore, changing the parameter value has no effect on the original value
	s = "Original value"
	t := s
	changeValue(s)
	fmt.Printf("s=%v (%T)\n", s, s)
	fmt.Printf("t=%v (%T)\n", t, t)
	// To change the original value, use a pointer
	sptr = &s
	changePointer(sptr)
	fmt.Printf("s=%v (%T)\n", s, s)
	fmt.Printf("t=%v (%T)\n", t, t)
}

// Utility function to obtain a pointer to a primitive
// whether it be a variable or a constant, because either
// way the value is copied to a parameter (variable)
func stringp(s string) *string {
	return &s
}

func changeValue(s string) {
	s = "New value"
}

func changePointer(s *string) {
	*s = "New value"
}

