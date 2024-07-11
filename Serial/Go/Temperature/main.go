package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"go.bug.st/serial"
)

var x float64
var y float64
var z float64

func main() {
	xposl := ""
	xposh := ""
	fmt.Println("Test Multiport Serial Controller")

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No Serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	mctl := 0
	po := false
	pos := 0
	for x := 0; x < len(ports); x++ {
		port, err := serial.Open(ports[x], mode)
		if err != nil {
			fmt.Println(err)
			po = false
		} else {
			po = true
		}
		if po {
			buff := make([]byte, 1)
			for {
				n, err := port.Read(buff)
				if err != nil {
					log.Fatal(err)
				}
				if n == 0 {
					port.Close()
					break
				}
				src := []byte(string(buff))

				encodedStr := hex.EncodeToString(src)
				if encodedStr == "55" {
					mctl = 1
					pos = 0
				}
				if encodedStr == "51" && mctl == 1 {
					fmt.Println("")
					mctl = 2
				}

				if mctl == 2 {
					pos++

					if err != nil {
						fmt.Println(err)
					}

					// fmt.Printf(" %s %d %d\n", encodedStr, decimal, pos)

					switch {
					case pos == 8:
						xposl = encodedStr
					case pos == 9:
						xposh = encodedStr
						byte1, _ := strconv.ParseUint(xposl, 16, 8)
						byte2, _ := strconv.ParseUint(xposh, 16, 8)
						combinedx := (uint16(byte1) << 8) | uint16(byte2)
						fmt.Printf("Combined hex fo X: 0x%X\n", combinedx)
						var s string = strconv.FormatUint(uint64(combinedx), 10)
						fmt.Printf("s=%s\n", s)
						f, err := strconv.ParseFloat(s, 32)
						if err == nil {
							fmt.Println(f) // 3.1415927410125732
						}
						fmt.Printf("s=%f\n", f/32768*96.38+36.53)

					}
				}
			}
		}
	}

}

func convertAndCombine(s1, s2 string) string {
	hex1 := hex.EncodeToString([]byte(s1))
	hex2 := hex.EncodeToString([]byte(s2))

	return hex1 + hex2
}
