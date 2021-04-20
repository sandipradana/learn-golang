package main

import "fmt"

func main() {

	// signed
	var int8 int8 = 127
	var int16 int16 = 32767
	var int32 int32 = 2147483647
	var int64 int64 = 9223372036854775807

	fmt.Println("signed", int8, int16, int32, int64)

	// unsigned
	var uint8 uint8 = 255
	var uint16 uint16 = 65535
	var uint32 uint32 = 4294967295
	var uint64 uint64 = 18446744073709551615

	fmt.Println("unsigned", uint8, uint16, uint32, uint64)

	// alias
	var byte byte = 127        // int 8
	var rune rune = 2147483647 // int32
	var int int = 2147483647   // int32

	fmt.Println("alias", byte, int, rune)
}
