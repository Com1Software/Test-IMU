
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"

	"go.bug.st/serial"
)

var xpos float64
var ypos float64

func main() {
	// ... (rest of your code)

	if mctl == 2 {
		pos++

		decimal, err := strconv.ParseInt(encodedStr, 16, 32)
		if err != nil {
			fmt.Println(err)
		}

		switch {
		case pos == 2:
			xpos = float64(decimal) / 32768.0 * 180.0
			fmt.Printf(" x= %f ", xpos)
			fmt.Printf(" xH= %d", decimal)
		case pos == 3:
			ypos = float64(decimal) / 32768.0 * 180.0
			fmt.Printf(" y= %f ", ypos)
			fmt.Printf(" yH= %d", decimal)
		// ... (rest of your switch cases)
		}
	}

	// Calculate the heading
	heading := math.Atan2(ypos, xpos) * 180 / math.Pi
	if heading < 0 {
		heading += 360
	}
	fmt.Printf("Heading: %f\n", heading)

	// ... (rest of your code)
}
