package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Go Standard Library: sql")
	fmt.Println()

	// For now, just trying out some NullXXX types...

	i0 := sql.NullInt64{}
	fmt.Printf("i0: %+v\n", i0)
	fmt.Printf("i0.Value: %d\n", i0.Int64)
	fmt.Println()

	var i1 sql.NullInt64 = sql.NullInt64{Int64: 123, Valid: true}
	fmt.Printf("i1: %+v\n", i1)
	fmt.Printf("i1.Value: %d\n", i1.Int64)
	fmt.Println()

	s0 := sql.NullString{}
	fmt.Printf("s0: %+v\n", s0)
	fmt.Printf("s0.Value: \"%s\"\n", s0.String)
	fmt.Println()

	var s1 sql.NullString = sql.NullString{String: "Hello, World!", Valid: true}
	fmt.Printf("s1: %+v\n", s1)
	fmt.Printf("s1.Value: \"%s\"\n", s1.String)
	fmt.Println()

	fmt.Println()
}
