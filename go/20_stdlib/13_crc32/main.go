package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	fmt.Println("Go Standard Library: crc32")
	fmt.Println()

	//s := "https://preview6.kaseya.net"
	//s := "https://central.pcsupportgroup.com"
	s := "https://support.myitsupport.co.za"
	checksum := crc32.ChecksumIEEE([]byte(s))
	fmt.Println("checksum:", checksum)
	fmt.Println()
}
