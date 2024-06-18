package main

import (
	"fmt"
	"time"
)

const (
	AX = iota
	GX
	Roll
	HX
)

var sReg [4][3]int16

func main() {
	var a, w, Angle, h [3]float32
	fmt.Print("Please enter the serial number: ")
	var iComPort int
	fmt.Scanln(&iComPort)

	OpenCOMDevice(iComPort, 9600)
	WitInit(0x50, WIT_PROTOCOL_NORMAL)
	WitSerialWriteRegister(SensorUartSend)
	WitRegisterCallBack(CopeSensorData)
	AutoScanSensor()

	for {
		time.Sleep(500 * time.Millisecond)
		for i := 0; i < 3; i++ {
			a[i] = float32(sReg[AX][i]) / 32768.0 * 16.0
			w[i] = float32(sReg[GX][i]) / 32768.0 * 2000.0
			Angle[i] = float32(sReg[Roll][i]) / 32768.0 * 180.0
			h[i] = float32(sReg[HX][i])
		}
		fmt.Printf("a:%.2f %.2f %.2f\r\n", a[0], a[1], a[2])
		fmt.Printf("w:%.2f %.2f %.2f\r\n", w[0], w[1], w[2])
		fmt.Printf("Angle:%.1f %.1f %.1f\r\n", Angle[0], Angle[1], Angle[2])
		fmt.Printf("h:%.0f %.0f %.0f\r\n\r\n", h[0], h[1], h[2])
	}
}

func DelayMs(ms uint16) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func SensorUartSend(p_data []byte, uiSize uint32) {
	SendUARTMessageLength(string(p_data), uiSize)
}

func CopeSensorData(uiReg, uiRegNum uint32) {
	s_cDataUpdate = 1
}

func AutoScanSensor() {
	c_uiBaud := []int{4800, 9600, 19200, 38400, 57600, 115200, 230400}
	var iRetry int

	for _, baud := range c_uiBaud {
		CloseCOMDevice()
		OpenCOMDevice(iComPort, baud)
		iRetry = 2
		for iRetry > 0 {
			s_cDataUpdate = 0
			WitReadReg(AX, 3)
			DelayMs(100)
			if s_cDataUpdate != 0 {
				fmt.Printf("%d baud find sensor\r\n\r\n", baud)
				return
			}
			iRetry--
		}
	}
	fmt.Println("can not find sensor")
	fmt.Println("please check your connection")
}

