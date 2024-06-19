package main

import "fmt"

func main() {
    var hex1 byte = 0x12 // First hex byte
    var hex2 byte = 0x34 // Second hex byte

    combined := (uint16(hex1) << 8) | uint16(hex2)

    fmt.Printf("Combined hex: 0x%X\n", combined)
}
