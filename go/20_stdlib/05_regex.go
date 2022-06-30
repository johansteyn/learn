package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Go Standard Library: regex")
	fmt.Println()

	s1 := "The quick brown fox."
	s2 := "The duck said quack!"
	s3 := "The quick platypus said quack, queck and quock..."
	s4 := "What is a quark?"
	pattern := "q[aeiou]+ck"

	// Using a regex pattern directly
  match, _ := regexp.MatchString(pattern, s1)
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s1, match)
  match, _ = regexp.MatchString(pattern, s2)
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s2, match)
  match, _ = regexp.MatchString(pattern, s3)
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s3, match)
  match, _ = regexp.MatchString(pattern, s4)
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s4, match)
	fmt.Println()

	// Compiling the regex pattern beforehand
  r, _ := regexp.Compile(pattern)
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s1, r.MatchString(s1))
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s2, r.MatchString(s2))
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s3, r.MatchString(s3))
  fmt.Printf("Pattern '%s' matches '%s'? %t\n", pattern, s4, r.MatchString(s4))
	fmt.Println()

	// Finds the part that matches (instead of just returning true/false)
  fmt.Printf("Pattern '%s' matches '%s' in '%s'\n", pattern, r.FindString(s1), s1)
  fmt.Printf("Pattern '%s' matches '%s' in '%s'\n", pattern, r.FindString(s2), s2)
  fmt.Printf("Pattern '%s' matches '%s' in '%s'\n", pattern, r.FindString(s3), s3)
  fmt.Printf("Pattern '%s' matches '%s' in '%s'\n", pattern, r.FindString(s4), s4)
	fmt.Println()

	// Finds the indexes of the matches
  fmt.Printf("Pattern '%s' matches '%s' at indexes: %v\n", pattern, s1, r.FindStringIndex(s1))
  fmt.Printf("Pattern '%s' matches '%s' at indexes: %v\n", pattern, s2, r.FindStringIndex(s2))
  fmt.Printf("Pattern '%s' matches '%s' at indexes: %v\n", pattern, s3, r.FindStringIndex(s3))
  fmt.Printf("Pattern '%s' matches '%s' at indexes: %v\n", pattern, s4, r.FindStringIndex(s4))
	fmt.Println()

	// Finds all parts that match
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s1, -1), s1)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s2, -1), s2)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s3, -1), s3)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s4, -1), s4)
	fmt.Println()
	
	// Specifying -1 above returns all matches
	// A non-negative value will limit the number of matches
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s1, 1), s1)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s2, 1), s2)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s3, 1), s3)
  fmt.Printf("Pattern '%s' matches %v in '%s'\n", pattern, r.FindAllString(s4, 1), s4)
	fmt.Println()

}

