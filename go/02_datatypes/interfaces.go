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

}

type shape interface {
	area() int // Note that there is no "func" keyword
}

type rectangle struct {
	width, height int
}

func (r rectangle) area() int {
	return r.width * r.height
}

type circle struct {
	radius int
}

func (c circle) area() int {
	return int(float32(c.radius * c.radius) * math.Pi)
}

