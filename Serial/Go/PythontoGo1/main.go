package main

import (
	"fmt"
	"time"
	"math"
	"encoding/hex"
)

func main() {
	// Initial conditions
	ser, err := serial.Open("/dev/ttyUSB0", 9600, 8, 'N', 1)
	if err != nil {
		fmt.Println("Error opening serial port:", err)
		return
	}
	defer ser.Close()

	readData := ""
	dataStartRecording := false

	// This part is needed so that the reading can start reliably.
	fmt.Println("Starting...", ser.Name())
	time.Sleep(1 * time.Second)
	ser.ResetInputBuffer()

	// Loop through the string of bytes. If you are planning to use this code in your project, the main loop starts here.
	for {
		rawData, err := ser.Read(2)
		if err != nil {
			fmt.Println("Error reading from serial port:", err)
			return
		}
		rawDataHex := hex.EncodeToString(rawData)

		// Make sure that ALL readData starts with 5551 before recording starts. Readings do not always start at bytes 5551 for some reason.
		if rawDataHex == "5551" {
			dataStartRecording = true
		}

		// Recording and concatenation functions
		if dataStartRecording {
			readData += rawDataHex
		}

		// Processing. Variable names based on variables on the Witmotion WT901CTTL datasheet.
		if len(readData) == 88 {
			StartAddress_1, _ := strconv.ParseInt(readData[0:2], 16, 64)
			StartAddress_A, _ := strconv.ParseInt(readData[2:4], 16, 64)
			AxL, _ := strconv.ParseInt(readData[4:6], 16, 64)
			AxH, _ := strconv.ParseInt(readData[6:8], 16, 64)
			AyL, _ := strconv.ParseInt(readData[8:10], 16, 64)
			AyH, _ := strconv.ParseInt(readData[10:12], 16, 64)
			AzL, _ := strconv.ParseInt(readData[12:14], 16, 64)
			AzH, _ := strconv.ParseInt(readData[14:16], 16, 64)
			TL_A, _ := strconv.ParseInt(readData[16:18], 16, 64)
			TH_A, _ := strconv.ParseInt(readData[18:20], 16, 64)
			SUM_A, _ := strconv.ParseInt(readData[20:22], 16, 64)

			StartAddress_2, _ := strconv.ParseInt(readData[22:24], 16, 64)
			StartAddress_w, _ := strconv.ParseInt(readData[24:26], 16, 64)
			wxL, _ := strconv.ParseInt(readData[26:28], 16, 64)
			wxH, _ := strconv.ParseInt(readData[28:30], 16, 64)
			wyL, _ := strconv.ParseInt(readData[30:32], 16, 64)
			wyH, _ := strconv.ParseInt(readData[32:34], 16, 64)
			wzL, _ := strconv.ParseInt(readData[34:36], 16, 64)
			wzH, _ := strconv.ParseInt(readData[36:38], 16, 64)
			TL_w, _ := strconv.ParseInt(readData[38:40], 16, 64)
			TH_w, _ := strconv.ParseInt(readData[40:42], 16, 64)
			SUM_w, _ := strconv.ParseInt(readData[42:44], 16, 64)

			StartAddress_3, _ := strconv.ParseInt(readData[44:46], 16, 64)
			StartAddress_ypr, _ := strconv.ParseInt(readData[46:48], 16, 64)
			RollL, _ := strconv.ParseInt(readData[48:50], 16, 64)
			RollH, _ := strconv.ParseInt(readData[50:52], 16, 64)
			PitchL, _ := strconv.ParseInt(readData[52:54], 16, 64)
			PitchH, _ := strconv.ParseInt(readData[54:56], 16, 64)
			YawL, _ := strconv.ParseInt(readData[56:58], 16, 64)
			YawH, _ := strconv.ParseInt(readData[58:60], 16, 64)
			VL, _ := strconv.ParseInt(readData[60:62], 16, 64)
			VH, _ := strconv.ParseInt(readData[62:64], 16, 64)
			SUM_ypr, _ := strconv.ParseInt(readData[64:66], 16, 64)

			StartAddress_4, _ := strconv.ParseInt(readData[66:68], 16, 64)
			StartAddress_mag, _ := strconv.ParseInt(readData[68:70], 16, 64)
			HxL, _ := strconv.ParseInt(readData[70:72], 16, 64)
			HxH, _ := strconv.ParseInt(readData[72:74], 16, 64)
			HyL, _ := strconv.ParseInt(readData[74:76], 16, 64)
			HyH, _ := strconv.ParseInt(readData[76:78], 16, 64)
			HzL, _ := strconv.ParseInt(readData[78:80], 16, 64)
			HzH, _ := strconv.ParseInt(readData[80:82], 16, 64)
			TL_mag, _ := strconv.ParseInt(readData[82:84], 16, 64)
			TH_mag, _ := strconv.ParseInt(readData[84:86], 16, 64)
			SUM_mag, _ := strconv.ParseInt(readData[86:88], 16, 64)

			// Acceleration output
			Ax := float64(int16((AxH<<8)|AxL))/32768.0*16.0
			Ay := float64(int16((AyH<<8)|AyL))/32768.0*16.0
			Az := float64(int16((AzH<<8)|AzL))/32768.0*16.0
			T_A := float64(int16((TH_A<<8)|TL_A))/100.0

			// Angular velocity output
			Wx := float64(int16((wxH<<8)|wxL))/32768.0*2000.0
			Wy := float64(int16((wyH<<8)|wyL))/32768.0*2000.0
			Wz := float64(int16((wzH<<8)|wzL))/32768.0*2000.0
			T_w := float64(int16((TH_w<<8)|TL_w))/100.0

			// Angle output
			Roll := float64(int16((RollH<<8)|RollL))/32768.0*180.0
			Pitch := float64(int16((PitchH<<8)|PitchL))/32768.0*180.0
			Yaw := float64(int16((YawH<<8)|YawL))/32768.0*180.0

			// Magnetic output
			Hx := float64(int16(HxH<<8)|HxL)
			Hy := float64(int16(HyH<<8)|HyL)
			Hz := float64(int16(HzH<<8)|HzL)
			T_mag := float64(int16((TH_mag<<8)|TL_mag))/100.0

			// Readable outputs. Uncomment for specific readouts.
			// fmt.Println(readData)
			// fmt.Printf("%6.3f %6.3f %6.3f\n", Ax, Ay, Az)
			fmt.Printf("%7.3f %7.3f %7.3f\n", Wx, Wy, Wz) // This detects any movement on the axes.
			// fmt.Printf("%7.3f\n", Roll)
		}
	}
}

