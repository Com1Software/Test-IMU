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
	yposl := ""
	yposh := ""
	zposl := ""
	zposh := ""
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
				if encodedStr == "53" && mctl == 1 {
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
					case pos == 2:
						xposl = encodedStr
					case pos == 3:
						xposh = encodedStr
						byte1, _ := strconv.ParseUint(xposl, 16, 8)
						byte2, _ := strconv.ParseUint(xposh, 16, 8)
						combinedx := (uint16(byte1) << 8) | uint16(byte2)/32768.0*180.0
						fmt.Printf("Combined hex fo X: 0x%X\n", combinedx)
						// var s string = strconv.FormatUint(uint64(combinedx), 10)
						//fmt.Printf("s=%s\n", s)

					case pos == 4:
						yposl = encodedStr
					case pos == 5:
						yposh = encodedStr
						byte1, _ := strconv.ParseUint(yposl, 16, 8)
						byte2, _ := strconv.ParseUint(yposh, 16, 8)
						combinedy := (uint16(byte1) << 8) | uint16(byte2)/32768.0*180.0
						fmt.Printf("Combined hex for Y: 0x%X\n", combinedy)

					case pos == 6:
						zposl = encodedStr
					case pos == 7:
						yposh = encodedStr
						byte1, _ := strconv.ParseUint(zposl, 16, 8)
						byte2, _ := strconv.ParseUint(zposh, 16, 8)
						combinedz := (uint16(byte1) << 8) | uint16(byte2)/32768.0*180.0
						fmt.Printf("Combined hex for Z: 0x%X\n", combinedz)

						//Roll = float(np.short((RollH<<8)|RollL)/32768.0*180.0)
						//Pitch = float(np.short((PitchH<<8)|PitchL)/32768.0*180.0)
						//Yaw = float(np.short((YawH<<8)|YawL)/32768.0*180.0)

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
