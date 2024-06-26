package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	"go.bug.st/serial"
)

var xposl float64
var xposh float64

func main() {
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

					decimal, err := strconv.ParseInt(encodedStr, 16, 32)
					if err != nil {
						fmt.Println(err)
					}

					// fmt.Printf(" %s %d %d\n", encodedStr, decimal, pos)

					switch {
					case pos == 2:
						xposl = float64(decimal) / 32768.0 * 180.0
						fmt.Printf(" xL= %f ", xposl)
						fmt.Printf(" xL= %d", decimal)
					case pos == 3:
						xposh = float64(decimal) / 32768.0 * 180.0
						fmt.Printf(" xH= %f ", xposh)
						fmt.Printf(" xH= %d", decimal)
					case pos == 4:
						fmt.Printf(" yH= %d", decimal)
					case pos == 5:
						fmt.Printf(" yL= %d", decimal)
					case pos == 6:
						fmt.Printf(" zH= %d", decimal)
					case pos == 7:
						fmt.Printf(" zL= %d", decimal)

					}
				}
			}
		}
	}

}
