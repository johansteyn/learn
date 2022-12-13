package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Standard Library: datestamp")
	fmt.Println()

  now := time.Now()
  fmt.Println(now.Format("2006-02-01")) // YYYY-DD-MM, ie. 2nd of January 2006
  fmt.Println(now.Format("2006-01-02")) // YYYY-MM-DD, ie. 2nd of January 2006
  fmt.Println(now.Format("20060102")) // YYYYMMDD, ie. 2nd of January 2006
}

