package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"regexp"

	"github.com/tarm/serial"
)

const (
	AngInit     = "\xff\xaa\x52"
	AccCalib    = "\xff\xaa\x67"
	declination = -0.00669
	pi          = 3.14159265359
	feature     = "UQ(.{6,6}).{3,3}UR(.{6,6}).{3,3}US(.{6,6}).{3,3}"
)

var (
	s_x, s_y, s_z float64
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	s.Write([]byte(AngInit))
	s.Write([]byte(AccCalib))

	for {
		buf := make([]byte, 65)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		r, _ := regexp.Compile(feature)
		match := r.Find(buf[:n])
		if match != nil {
			var frame [33]byte
			binary.Read(bytes.NewReader(match), binary.LittleEndian, &frame)

			var sum_Q, sum_R, sum_S byte
			for i := 0; i < 10; i++ {
				sum_Q += frame[i]
				sum_R += frame[i+11]
				sum_S += frame[i+22]
			}

			if sum_Q == frame[10] && sum_R == frame[21] && sum_S == frame[32] {
				var af, wf, ef [3]int16
				binary.Read(bytes.NewReader(match[1:7]), binary.LittleEndian, &af)
				binary.Read(bytes.NewReader(match[14:20]), binary.LittleEndian, &wf)
				binary.Read(bytes.NewReader(match[27:33]), binary.LittleEndian, &ef)

				var af_l, wf_l, ef_l [3]float64
				for i := 0; i < 3; i++ {
					af_l[i] = float64(af[i]) / 32768.0 * 16 * 9.8
					wf_l[i] = float64(wf[i]) / 32768.0 * 2000
					ef_l[i] = float64(ef[i]) / 32768.0 * 180
				}

				s_x += af_l[0] * 0.01
				s_y += af_l[1] * 0.01
				s_z += (af_l[2] - 9.8) * 0.01

				heading := math.Atan2(float64(ef_l[1]), float64(ef_l[0])) + declination
				if heading > 2*pi {
					heading -= 2 * pi
				}
				if heading < 0 {
					heading += 2 * pi
				}

				fmt.Println("--- angle ---")
				fmt.Println(ef_l[0], ef_l[1], ef_l[2])
				heading_angle := int(heading * 180 / pi)
				fmt.Printf("Heading Angle = %d°\n", heading_angle)
			}
		}
	}
}
