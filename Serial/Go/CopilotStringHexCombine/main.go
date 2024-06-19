package main

import (
	"encoding/hex"
	"fmt"
)

func convertAndCombine(s1, s2 string) string {
	hex1 := hex.EncodeToString([]byte(s1))
	hex2 := hex.EncodeToString([]byte(s2))

	return hex1 + hex2
}

func main() {
	s1 := "Hello"
	s2 := "World"
	fmt.Println(convertAndCombine(s1, s2))
}
