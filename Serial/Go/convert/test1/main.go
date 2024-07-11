package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
	"unsafe"
)

var (
	AngInit     = []byte{0xff, 0xaa, 0x52}
	AccCalib    = []byte{0xff, 0xaa, 0x67}
	declination = -0.00669
	pi          = 3.14159265359
	feature     = `UQ(.{6,6}).{3,3}UR(.{6,6}).{3,3}US(.{6,6}).{3,3}`
	fmt_B       = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	fmt_h       = "<hhh"
	s_x, s_y, s_z float64
)

func main() {
	ser, err := serial.Open("/dev/ttyUSB0", 115200)
	if err != nil {
		fmt.Println("failed to open serial port")
		return
	}
	defer ser.Close()

	fmt.Println("success")

	ser.Write(AngInit)
	ser.Write(AccCalib)

	for {
		imu_msg, err := ser.Read(65)
		if err != nil {
			fmt.Println("error reading from serial port:", err)
			continue
		}

		result := regexp.MustCompile(feature).FindSubmatch(imu_msg)
		if result != nil {
			frame := make([]byte, len(result[0]))
			copy(frame, result[0])

			hex_string := fmt.Sprintf("% x", frame)
			sum_Q, sum_R, sum_S := 0, 0, 0
			for i := 0; i < 10; i++ {
				sum_Q += int(frame[i])
				sum_R += int(frame[i+11])
				sum_S += int(frame[i+22])
			}
			sum_Q &= 0x000000ff
			sum_R &= 0x000000ff
			sum_S &= 0x000000ff

			if sum_Q == int(frame[10]) && sum_R == int(frame[21]) && sum_S == int(frame[32]) {
				af := *(*[3]int16)(unsafe.Pointer(&result[1][0]))
				wf := *(*[3]int16)(unsafe.Pointer(&result[2][0]))
				ef := *(*[3]int16)(unsafe.Pointer(&result[3][0]))

				af_l := make([]float64, 3)
				wf_l := make([]float64, 3)
				ef_l := make([]float64, 3)
				for i := 0; i < 3; i++ {
					af_l[i] = round(float64(af[i])/32768.0*16, 2) * 9.8
					wf_l[i] = float64(wf[i]) / 32768.0 * 2000
					ef_l[i] = round(float64(ef[i])/32768.0*180, 2)
				}
				linear_acceleration_x, linear_acceleration_y, linear_acceleration_z := af_l[0], af_l[1], af_l[2]

				s_x += linear_acceleration_x * 0.01
				s_y += linear_acceleration_y * 0.01
				s_z += (linear_acceleration_z - 9.8) * 0.01

				angular_velocity_x, angular_velocity_y, angular_velocity_z := wf_l[0], wf_l[1], wf_l[2]

				roll, pitch, yaw := ef_l[0], ef_l[1], ef_l[2]
				fmt.Println("--- angle ---")
				heading := math.Atan2(pitch, roll) + declination
				if heading > 2*pi {
					heading -= 2 * pi
				}
				if heading < 0 {
					heading += 2 * pi
				}
				heading_angle := int(heading * 180 / pi)
				fmt.Printf("Heading Angle = %d°\n", heading_angle)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func round(val float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	round = float64(int64(digit + math.Copysign(0.5, digit)))
	return round / pow
}

func serial.Open(port string, baud int) (*serial.Port, error) {
	// Implement serial port opening and configuration
	// Return a serial.Port instance and an error (if any)
}

func (p *serial.Port) Read(b []byte) (int, error) {
	// Implement serial port read operation
	// Return the number of bytes read and an error (if any)
}

func (p *serial.Port) Write(b []byte) (int, error) {
	// Implement serial port write operation
	// Return the number of bytes written and an error (if any)
}

func (p *serial.Port) Close() error {
	// Implement serial port closing operation
	// Return an error (if any)
}

