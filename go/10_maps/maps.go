package main

import (
  "fmt"
  "sort"
)

func main() {
  ages := map[string]int {
    "Bob": 54,
    "Carol": 42,
    "Alice": 31,
  }
	fmt.Println(ages)
  var names[]string
  for name := range ages {
    names = append(names, name)
  }
  sort.Strings(names)
	fmt.Println(names)
  for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
  }
}
