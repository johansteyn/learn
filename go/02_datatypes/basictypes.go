package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Basic Data Types")

	// Four kinds of declarations: var, const, type and func
	// The standard way to declare variables is with "var",
	// followed by the variable name, then the type.
	// A value may optionally be assigned.
	// If a value is assigned the the type is optional (it will be inferred)
	// If no value is assigned, the variable will get its default nil value (based on type)

	// Integers
	var i int // Default (nil) value is 0
	fmt.Printf("i=%d (%T)\n", i, i)
	var j int = 42  // Variable type follows variable name and is optional.
	fmt.Printf("j=%d (%T)\n", j, j)
	var k int = math.MaxInt64 // On my MacBook an int is 64 bits 
	fmt.Printf("k=%d (%T)\n", k, k)
	var i8 int8 = math.MaxInt8
	fmt.Printf("i8=%d (%T)\n", i8, i8)
	var i16 int16 = math.MaxInt16
	fmt.Printf("i16=%d (%T)\n", i16, i16)
	var i32 int32 = math.MaxInt32
	fmt.Printf("i32=%d (%T)\n", i32, i32)
	var i64 int64 = math.MaxInt64
	fmt.Printf("i64=%d (%T)\n", i64, i64)

	// Unsigned Integers
	var ui uint // Default (nil) value is 0
	fmt.Printf("ui=%d (%T)\n", ui, ui)
	var ui8 uint8 = math.MaxUint8
	fmt.Printf("ui8=%d (%T)\n", ui8, ui8)
	var ui16 uint16 = math.MaxUint16
	fmt.Printf("ui16=%d (%T)\n", ui16, ui16)
	var ui32 uint32 = math.MaxUint32
	fmt.Printf("ui32=%d (%T)\n", ui32, ui32)
	var ui64 uint64 = math.MaxUint64
	fmt.Printf("ui64=%d (%T)\n", ui64, ui64)
	var bite byte = math.MaxUint8 // The byte type is an alias for uint8
	fmt.Printf("bite=%d (%T)\n", bite, bite)

	// Default (inferred) type is int, which can be 32 or 64 bit depending on platform.
	var million = 1_000_000 // Underscores help with readability (but don't do something like 1_0_0_0_0_0_0)
	fmt.Printf("million=%d (%T)\n", million, million)

	// Short declaration form - the "var" keyword and type are omitted and a value is assigned with :=
	// Note that := is a declaration, whereas = is an assignment (but see exception below)
	n := 1_234_567
	fmt.Printf("n=%d (%T)\n", n, n)
	// The "new" function can be used to assign the default nil value with the short declaration form
	// But note that it returns a pointer...
	o := new(int)
	fmt.Printf("o=%d (%T)\n", *o, *o)
	// Note that the short declaration form does not work outside of functions
	// For idiomatic Go, avoid the short declaration form:
	// - When initializing a variable to its zero value
	// - To avoid explicit conversions, eg:
	//   x := byte(20)		<= legal but requires explicit conversion
	//   var x byte = 20	<= no explicit conversion
	// - To avoid unintended "shadowing", since it allows you to assign to both new and existing variables (see further down)
	//   By using standard variable declaration you are making it clear that this is a new variable

	// Binary, Octal and Hexadecimal integer literals
	ib := 0b101010 // Binary 42
	fmt.Printf("ib=%d (%T)\n", ib, ib)
	io := 0o52 // Octal 42 (Can also use 052, but not recommended)
	fmt.Printf("io=%d (%T)\n", io, io)
	ih := 0x2A // Hexadecimal 42
	fmt.Printf("ih=%d (%T)\n", ih, ih)

	// Float
	var f float32 // Default nil value is 0.0
	fmt.Printf("f=%f (%T)\n", f, f)
	var f32 float32 = 6.03E23 // Can use lowercase 'e' or uppercase 'E'
	fmt.Printf("f32=%f (%T)\n", f32, f32)
	var f64 float64 = 6.03E23
	fmt.Printf("f64=%f (%T)\n", f64, f64)

	// Complex numbers (TODO...)

	// Boolean
	var b bool // Default nil value is false
	fmt.Printf("b=%t (%T)\n", b, b)
	bt := true
	fmt.Printf("bt=%t (%T)\n", bt, bt)
	bf := false
	fmt.Printf("bf=%t (%T)\n", bf, bf)

	// Rune
	// A rune is an integer value that represents a character
	// The rune type is an alias for int32
	// Rune literals are surrounded by single quotes
	var r rune // Default nil value is 0, which displays as an empty character
	fmt.Printf("r=%c (%T)\n", r, r)
	a := 'a'
	fmt.Printf("a=%c (%T)\n", a, a)
	ad := 97 // Decimal
	fmt.Printf("ad=%c (%T)\n", ad, ad)
	ao := '\141' // Octal
	fmt.Printf("ao=%c (%T)\n", ao, ao)
	ah := '\x61' // Hexadecimal
	fmt.Printf("ah=%c (%T)\n", ah, ah)
	au16 := '\u0061' // 16-bit Unicode
	fmt.Printf("au16=%c (%T)\n", au16, au16)
	au32 := '\U00000061' // 32-bit Unicode - note the uppercase U
	fmt.Printf("au32=%c (%T)\n", au32, au32)

	// Since runes are integers, we can use integer operators
	e := a + 4
	fmt.Printf("e=%c (%T)\n", e, e)
	d := e - 1
	fmt.Printf("d=%c (%T)\n", d, d)

	// Any UTF-8 character can be used directly
	pi := 'π'
	fmt.Printf("pi=%c (%T)\n", pi, pi)

	// Escaped runes
	newline := '\n'
	fmt.Printf("newline=%c (%T)\n", newline, newline)
	tab := '\t'
	fmt.Printf("tab=%c (%T)\n", tab, tab)
	singleQuote := '\''
	fmt.Printf("singleQuote=%c (%T)\n", singleQuote, singleQuote)
	// Page 18 of "Learning Go" says the double quote can be escaped, but it can't (and why would you?)
	doubleQuote := '"'
	fmt.Printf("doubelQuote=%c (%T)\n", doubleQuote, doubleQuote)
	backslash := '\\'
	fmt.Printf("backslash=%c (%T)\n", backslash, backslash)

	// Multiple variables can be declared in a single statement
	var width, height int = 100, 200
	fmt.Printf("width=%d (%T)\n", width, width)
	fmt.Printf("height=%d (%T)\n", height, height)
	// The type is optional, and they may even be different types
	var name, age = "Johan", 42
	fmt.Printf("name=%s (%T)\n", name, name)
	fmt.Printf("age=%d (%T)\n", age, age)
	// Also works with the short declaration form
	x, y := 123, 456
	fmt.Printf("x=%d (%T)\n", x, x)
	fmt.Printf("y=%d (%T)\n", y, y)
	// But be careful... at least one variable must be new, which means the others could "shadow" previous declarations
	// ie. the new variables are being declared, whereas the existing variables are assigned new values.
	// Therefore, in this case, := is used to both declare and assign.
	x, y, z := 987, 654, 321
	fmt.Printf("x=%d (%T)\n", x, x)
	fmt.Printf("y=%d (%T)\n", y, y)
	fmt.Printf("z=%d (%T)\n", z, z)

	// Another way is using a declaration list
	var (
		l1 int
		l2 int = 1234
		l3 = 5678
		l4 = "abc"
		l5, l6 = 42, "def"
	)
	fmt.Printf("l1=%d (%T)\n", l1, l1)
	fmt.Printf("l2=%d (%T)\n", l2, l2)
	fmt.Printf("l3=%d (%T)\n", l3, l3)
	fmt.Printf("l4=%s (%T)\n", l4, l4)
	fmt.Printf("l5=%d (%T)\n", l5, l5)
	fmt.Printf("l6=%s (%T)\n", l6, l6)

	// Variables of different types cannot be assigned to one another
	//i = f // Cannot assign a float to an int
	//f = i // Cannot even assign an int to a float
	//i32 = i16 // Nor even an int16 to an int32
	// In other words: Go does not have implicit conversion - all conversions have to be explicit.
	i32 = int32(i16) // Explicit conversion (with no loss)
	fmt.Printf("i32=%d (%T)\n", i32, i32)
	i16 = int16(i64) // Explicit conversion (with loss)
	fmt.Printf("i16=%d (%T)\n", i16, i16)

	// But note that literals (and constants) are untyped, which allows you to assign what looks like an int literal to a float
	f = 123456
	fmt.Printf("f=%f (%T)\n", f, f)
	// Yet you still can't assign a float literal to an int
	//i = 6.03e23 // Implicit conversion fails
	//i = int(6.03e23) // Explicit conversion also fails (because the literal value overflows int)
	// Even a smaller float literal that doesn't overflow cannot be assigned, since it results in truncation 
	//i = 1.23
	//i = int(1.23)

	// Constants
	// A constant in Go is simply a way to give a name to a literal value
	const c = 12
	fmt.Printf("c=%d (%T)\n", c, c)
	//c = 7 // Cannot re-assign a value to a constant

	// Literals (and therefore constants) are untyped
	//   https://riptutorial.com/go/example/12431/typed-vs--untyped-constants
	// While literals/constants have no types of their own, they have a default type when np other type can be inferred.
	// A typed constant ensures that only a value of that type can be assigned to it,
	// whereas leaving a constant untyped gives more flexibility

	// So, our constant c above is untyped, which means it can be assigned to all these variables:
  var ci int = c
	fmt.Printf("ci=%d (%T)\n", ci, ci)
  var cf float32 = c
	fmt.Printf("cf=%f (%T)\n", cf, cf)
  var cb byte = c
	fmt.Printf("cb=%d (%T)\n", cb, cb)

	// A typed constant
	const tc int = 24
	fmt.Printf("tc=%d (%T)\n", tc, tc)
  var tci int = tc // Can assign the typed constant to an int variable
	fmt.Printf("tci=%d (%T)\n", tci, tci)
  //var tcf float32 = tc // Cannot assign the typed constant to any non-int variable
	//fmt.Printf("tcf=%f (%T)\n", tcf, tcf)

	// Unused variables are not allowed
	//var unusedVar int = 36
	// Unused constants are fine because constants cannot have side-effects, so can be eliminated by the compiler
	const unusedConst int = 48

	// NOTE: Go does not use ALL_CAPS convention for constant names.
	// This is mainly because all members that start with an uppercase letter are public.
	const MAX_TEMP = 100 // Don't use
	const maxTemp = 100  // Rather use camelCase

	// Types
	// A type declaration defines a new named type for an existing underlying type
	// Form: type <name> <underlying-type>
	// Types are usually declared at package level (see below)
	var freezingC celsius = 0
	var freezingF fahrenheit = 32
	// Cannot compare values of different types, even if their underlying types are the same
	//if freezingC == freezingF {
	// But can do a type conversion if their underlying types are the same
	var convertedF = fahrenheit(freezingC)
	fmt.Printf("convertedF = %d\n", convertedF)
	if freezingF == convertedF {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", freezingF, freezingC)
	} else {
		fmt.Printf("%d Fahrenheit != %d Celsius\n", freezingF, freezingC)
	}
	// But that's not what we really want...
	// We want to call a proper conversion function that calculates the correct value
	var calculatedF = celsiusToFahrenheit(freezingC)
	fmt.Printf("calculatedF = %d\n", calculatedF)
	if freezingF == calculatedF {
		fmt.Printf("%d Fahrenheit == %d Celsius\n", freezingF, freezingC)
	} else {
		fmt.Printf("%d Celsius != %d Fahrenheit\n", freezingF, freezingC)
	}

}

type celsius int
type fahrenheit int
func celsiusToFahrenheit (c celsius) fahrenheit {
	return fahrenheit(c * 9 / 5 + 32)
}

