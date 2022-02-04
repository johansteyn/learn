package main

import (
	"fmt"
)

func main() {
	fmt.Println("Operators")
	fmt.Println()

	// Arithmetic
	x, y, z := 12, 3, 7
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
	fmt.Printf("%d - %d = %d\n", x, y, x - y)
	fmt.Printf("%d * %d = %d\n", x, y, x * y)
	fmt.Printf("%d / %d = %d\n", x, y, x / y)
	fmt.Printf("%d %% %d = %d\n", x, z, x % z) // Note how % is escaped
	fmt.Println()

	// Unary arithmetic
	fmt.Printf("-%d = %d\n", x, -x)
	fmt.Printf("+%d = %d\n", x, +x)
	fmt.Println()

	// Increment/decrement and assignment
	n := 7
	fmt.Printf("n: %d\n", n)
	n++
	fmt.Printf("n++: %d\n", n)
	n--
	fmt.Printf("n--: %d\n", n)
	n += 5
	fmt.Printf("n += 5: %d\n", n)
	n -= 7
	fmt.Printf("n -= 7: %d\n", n)
	n *= 9
	fmt.Printf("n *= 9: %d\n", n)
	n /= 3
	fmt.Printf("n /= 3: %d\n", n)
	n %= 8
	fmt.Printf("n %%= 8: %d\n", n) // Back to where we were :)
	fmt.Println()

	// Bitwise
	a, b, c := 0, 6, 15 // 0000, 0110 and 1111
	fmt.Printf("%04b & %04b = %04b\n", a, b, a & b)
	fmt.Printf("%04b & %04b = %04b\n", b, c, b & c)
	fmt.Printf("%04b & %04b = %04b\n", a, c, a & c)
	fmt.Printf("%04b | %04b = %04b\n", a, b, a | b)
	fmt.Printf("%04b | %04b = %04b\n", b, c, b | c)
	fmt.Printf("%04b | %04b = %04b\n", a, c, a | c)
	fmt.Printf("%04b ^ %04b = %04b\n", a, b, a ^ b)
	fmt.Printf("%04b ^ %04b = %04b\n", b, c, b ^ c)
	fmt.Printf("%04b ^ %04b = %04b\n", a, c, a ^ c)
	// &^ is bit clear (AND NOT)
	fmt.Printf("%04b &^ %04b = %04b\n", a, b, a &^ b)
	fmt.Printf("%04b &^ %04b = %04b\n", b, c, b &^ c)
	fmt.Printf("%04b &^ %04b = %04b\n", a, c, a &^ c)
	fmt.Printf("%04b &^ %04b = %04b\n", b, a, b &^ a)
	fmt.Printf("%04b &^ %04b = %04b\n", c, b, c &^ b)
	fmt.Printf("%04b &^ %04b = %04b\n", c, a, c &^ a)
	// Shift
	fmt.Printf("%04b << %d = %04b (%d)\n", a, 1, a << 1, a << 1)
	fmt.Printf("%04b << %d = %04b (%d)\n", b, 1, b << 1, b << 1)
	fmt.Printf("%04b << %d = %04b (%d)\n", b, 2, b << 2, b << 2)
	fmt.Printf("%04b << %d = %04b (%d)\n", b, 3, b << 3, b << 3)
	fmt.Printf("%04b << %d = %04b (%d)\n", c, 1, c << 1, c << 1)
	fmt.Printf("%04b >> %d = %04b (%d)\n", a, 1, a >> 1, a >> 1)
	fmt.Printf("%04b >> %d = %04b (%d)\n", b, 1, b >> 1, b >> 1)
	fmt.Printf("%04b >> %d = %04b (%d)\n", b, 2, b >> 2, b >> 2)
	fmt.Printf("%04b >> %d = %04b (%d)\n", b, 3, b >> 3, b >> 3)
	fmt.Printf("%04b >> %d = %04b (%d)\n", c, 1, c >> 1, c >> 1)
	// Shifting by zero has no effect
	fmt.Printf("%04b << %d = %04b (%d)\n", b, 0, b << 0, b << 0)
	fmt.Printf("%04b >> %d = %04b (%d)\n", b, 0, b >> 0, b >> 0)
	// Cannot shift a negative amount of places
	// In previous versions of Go a negative value would be cast to an unsigned int,
	// which would work but with an unexpected result...
	// Since Go 1.13 specifying a negative number will simply result in a panic
	// (which is better that silently not working as expected...)
	//fmt.Printf("%04b << %d = %04b (%d)\n", a, 1, a << -1, a << -1)
	fmt.Println()

	// Logical
	t, f := true, false
	fmt.Printf("%t && %t = %t\n", f, f, f && f)
	fmt.Printf("%t && %t = %t\n", f, t, f && t)
	fmt.Printf("%t && %t = %t\n", t, f, t && f)
	fmt.Printf("%t && %t = %t\n", t, t, t && t)
	fmt.Printf("%t || %t = %t\n", f, f, f || f)
	fmt.Printf("%t || %t = %t\n", f, t, f || t)
	fmt.Printf("%t || %t = %t\n", t, f, t || f)
	fmt.Printf("%t || %t = %t\n", t, t, t || t)
	fmt.Println()

	// Unary logical
	fmt.Printf("!%t = %t\n", t, !t)
	fmt.Printf("!%t = %t\n", f, !f)
	fmt.Println()

	// Comparison
	fmt.Printf("%d + %d == %d? %t\n", x, y, z, x + y == z)
	fmt.Printf("%d + %d == %d? %t\n", x, y, c, x + y == c)
	fmt.Printf("%d + %d != %d? %t\n", x, y, z, x + y != z)
	fmt.Printf("%d + %d != %d? %t\n", x, y, c, x + y != c)
	fmt.Printf("%d < %d? %t\n", x, x, x < x)
	fmt.Printf("%d > %d? %t\n", x, x, x > x)
	fmt.Printf("%d <= %d? %t\n", x, x, x <= x)
	fmt.Printf("%d >= %d? %t\n", x, x, x >= x)
	fmt.Println()

	// Strings
	fox, dog := "The quick brown fox", "jumps over the lazy dog"
	fmt.Printf(`"%s" + "%s" = "%s"`, fox, dog, fox + dog)
	fmt.Println()
	fmt.Println()

	// Omitted: 
	//   Pointers (& and *)
	//   Channels (<-)
}

