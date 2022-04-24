package main

import (
	"fmt"
  "unicode/utf8"
)

func main() {
	fmt.Println("Strings")
	fmt.Println()

	var s string // Default nil value is an empty string
	fmt.Printf("s=%s (%T)\n", s, s)
	s = "String literals are delimited with \"double quotes\" (so may contain `backquotes`)\n and newlines must be escaped."
	fmt.Printf("s=%s (%T)\n", s, s)
	s = `Raw string literals are delimited by backquotes (so may contain "double quotes") 
but cannot contain backquotes at all
and newlines do not need to be escaped.`
	fmt.Printf("s=%s (%T)\n", s, s)
	s = "String literals contain runes: π a\142\x63\u0064\U00000065"
	fmt.Printf("s=%s (%T)\n", s, s)
	fmt.Println()

	// As with arrays and slices, you can access individual elements of a string
	quickbrownfox := "The quick brown fox jumps over the lazy dog"
	c := quickbrownfox[7]
	fmt.Printf("c=%c (%T)\n", c, c)
	x := quickbrownfox[18]
	fmt.Printf("x=%c (%T)\n", x, x)
	// But since strings are immutable, you cannot assign a value to an element
	//quickbrownfox[0] = 'Y'
	fmt.Println()

	// And you can slice a string
	fox := quickbrownfox[16:19]
	fmt.Printf("fox=%s (%T)\n", fox, fox)
	fmt.Println()

	// And iterate through the elements
	fmt.Print("quickbrownfox=")
	for i := 0; i < len(quickbrownfox); i++ {
		fmt.Printf("%c", quickbrownfox[i])
	}
	fmt.Println()
	// Also using range
	fmt.Print("quickbrownfox=")
	for _, c := range quickbrownfox {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	fmt.Println()

	// Up to now everything worked fine because we have been using UTF-8 codepoints that are 1 byte long
	// Things change when we have longer codepoints...
	foreign := "€£¥«»àèìòùëêôñ¿©®π§¹²³⁴±ΑΒΓΔΕ▶◀¼½¾█♂♀░♠♡♣♢《》〇〄【】✠〓✓●◐◑⇒♩♪♫あきㄤホ"
	fmt.Printf("foreign=%s (%T)\n", foreign, foreign)
	// There are only 62 characters, yet the string length is 156 (the length in bytes)
	fmt.Printf("len(foreign)=%d\n", len(foreign))
	// This function counts the number of characters
	fmt.Printf("count(foreign)=%d\n", count(foreign))
	// Rather use this standard library function though
	fmt.Printf("utf8.RuneCountInString(foreign)=%d\n", utf8.RuneCountInString(foreign))
	fmt.Println()

	// Accessing an individual element we only get the first byte
	euro := foreign[0]
	fmt.Printf("euro=%c (%T)\n", euro, euro)
	// Slicing also treats the string as a sequence of bytes
	currencies := foreign[:3]
	fmt.Printf("currencies=%s (%T)\n", currencies, currencies)
	// Similarly, iterating through the elements prints the wrong characters
	fmt.Print("foreign=")
	for i := 0; i < len(foreign); i++ {
		fmt.Printf("%c", foreign[i])
	}
	fmt.Println()
	// Though using range works fine
	fmt.Print("foreign=")
	for _, f := range foreign {
		fmt.Printf("%c", f)
	}
	fmt.Println()
	// Explanation: Go strings are not made out of runes - they are simply
	// a sequence of bytes, which don't have to be in any encoding.
	// However, several Go library functions, like "range", interpret the bytes as UTF-8 codepoints.
	// Conclusion: Only use "len", indexing and slicing on strings that contain only 1-byte codepoints.
	fmt.Println()

	// Type conversions between runes, strings and bytes
	// Create a string from a rune
	var r rune = 'a'
	fmt.Printf("r=%c (%T)\n", r, r)
	s = string(r)
	fmt.Printf("s=%s (%T)\n", s, s)
	// Create a string from a byte
	var b byte = 'b'
	fmt.Printf("b=%c (%T)\n", b, b)
	s = string(b)
	fmt.Printf("s=%s (%T)\n", s, s)
	// A string can also be created from an int, but don't (go vet will block it)
	var i int = 105 // ASCII 105 is the letter 'i'
	fmt.Printf("i=%c (%T)\n", i, i)
	s = string(i) // Many developers think this will result in string "105" instead of "i"
	fmt.Printf("s=%s (%T)\n", s, s)
	fmt.Println()

	// Convert a string to a slice of bytes
	s = "a€"
	fmt.Printf("s=%s (%T)\n", s, s)
	var bytes []byte = []byte(s)
	fmt.Println("bytes:", bytes)
	// Convert a string to a slice of runes
	var runes []rune = []rune(s)
	fmt.Print("runes: ")
	fmt.Println(runes)
}

func count (s string) int {
	var counter int
  for range s {
		counter++
  }
	return counter
}
