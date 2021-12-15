package main

import (
	"fmt"
)

func main() {
	fmt.Println("Flow Control")

	// If-Else is similar to other languages
	// No parentheses, and braces are mandatory
	if 1 + 1 == 2 {
		fmt.Println("One plus one equals two")
	} else {
		fmt.Println("What the heck...")
	}
	// Can declare a local variable that has scope in both the if and else blocks
	if x := 2; x + x == 4 {
		// x is in scope here
		fmt.Printf("%d + %d equals four\n", x, x)
	} else {
		// x is also in scope here
		fmt.Printf("%d + %d equals something other than four?!\n", x, x)
	}
	// x is out of scope here
	//fmt.Printf("x=%d\n", x)

	// Another way to do selection is with the switch statement
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, number := range numbers {
		// No parentheses, and braces are mandatory
		switch number {
		case 1, 2: // Multiple values allowed
			fmt.Printf("%d is a small number\n", number)
			// No break needed (ie. no fall-through) 
		case 7, 8, 9: {
			// Case clauses don't need braces, but they may be used
			fmt.Printf("%d is a large number\n", number)
		}
		default:
			// The default clause is optional
			fmt.Printf("%d is a medium number\n", number)
		}
		// Like if, we can declare a variable that has scope only in the switch block
		switch x := number % 2; x {
		case 0:
			fmt.Printf("%d is even (%d is zero)\n", number, x)
		case 1:
			fmt.Printf("%d is odd (%d is one)\n", number, x)
		default:
			// We can also declare a variable here
			y := 7
			fmt.Printf("y=%d)\n", y)
		}
	}
	// x and y are out of scope here
	//fmt.Printf("x=%d, y=%d\n", x, y)
	// You can switch on any values that are comparable (ie. not only integers)
	names := []string{"Alice", "Bob", "Carol", "Robin"}
	for _, name := range names {
		switch name {
		case "Alice", "Carol":
			fmt.Printf("%s is female\n", name)
		default:
			fmt.Printf("%s is male\n", name)
		}
	}



	// Go distinguishes between expression switches and type switches
	// TODO...

	// Loops
	// Four different ways - all using "for" statements
	// 1. C-style: initializer; conditional; incrementor
	for i := 0; i < 5; i+=2 {
		fmt.Printf("#%d\n", i)
	}
	fmt.Println()

	// 2. Condition-only
	i := 3
	for i > 0 {
		fmt.Printf("#%d\n", i)
		i--
	}
	fmt.Println()

	// 3. Infinite (along with continue and break)
	for {
		i++
		if i % 2 == 0 {
			continue;
		}
		if i % 7 == 0 {
			break;
		}
		fmt.Printf("#%d\n", i)
	}
	fmt.Println()

	// 4. Range

	// Goto
	// TODO...
}
