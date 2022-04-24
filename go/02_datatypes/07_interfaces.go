package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Interfaces")
	fmt.Println()

	rect := rectangle{
		width: 7,
		height: 9,
	}
	circ := circle{
		radius: 5,
	}
	shapes := []shape{rect, circ}
	for i, s := range(shapes) {
		fmt.Printf("Shape #%d has area = %d\n", i, s.area())
	}

	fmt.Printf("circ = %v\n", circ)

	// circ can be assigned to variables of either type
	var ss1 ShapeStringer1 = circ
	var ss2 ShapeStringer2 = circ
	fmt.Printf("ss1 = %v\n", ss1)
	fmt.Printf("ss2 = %v\n", ss2)
}

type shape interface {
	area() int // Note that there is no "func" keyword
}

type rectangle struct {
	width, height int
}

// Here rectangle implements the shape interface, but it does so implicityly
// ie. there is no explicit "implements" keyword
// This decouples interface from implementation - "duck typing" with type safety.
func (r rectangle) area() int {
	return r.width * r.height
}

type circle struct {
	radius int
}

func (c circle) area() int {
	return int(float32(c.radius * c.radius) * math.Pi)
}

// By convention, interface names often end with "er"
// Eg: Reader, Writer, Closer, Stringer...
type Stringer interface {
	String() string
}

func (c circle) String() string {
	return fmt.Sprintf("circle{%d}", c.radius)
}

// Since circle has both area and String methods, it implicitly
// implements both the Shape and Stringer interfaces, as well as
// the ShapeStringer interface, which can either contain both
// the area and String methods
type ShapeStringer1 interface {
	area() int
	String() string
}
// Or the Shape and Stringer interfaces can be embedded
type ShapeStringer2 interface {
	shape
	Stringer
}

