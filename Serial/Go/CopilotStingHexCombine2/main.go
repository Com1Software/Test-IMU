package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex1 := "12" // First hex byte as string
	hex2 := "34" // Second hex byte as string

	// Parse the hex strings to uint64, then convert to byte
	byte1, _ := strconv.ParseUint(hex1, 16, 8)
	byte2, _ := strconv.ParseUint(hex2, 16, 8)

	// Combine the bytes
	combined := (uint16(byte1) << 8) | uint16(byte2)

	fmt.Printf("Combined hex: 0x%X\n", combined)
}
