package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/paypal/gatt"
)

var msgsNum int

func onPeriphConnected(peripheral gatt.Peripheral, err error) {
	defer peripheral.Device().CancelConnection(peripheral)

	if err != nil {
		log.Fatalf("Failed to connect to peripheral: %s", err)
		return
	}

	fmt.Println("Connected to ", peripheral.ID())

	// ... Here you should discover services, characteristics etc.
	// to interact with the device.
}

func onPeriphDisconnected(peripheral gatt.Peripheral, err error) {
	fmt.Println("Disconnected")
}

func main() {
	device, err := gatt.NewDevice()
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
		return
	}

	// Register handlers.
	device.Handle(
		gatt.PeripheralConnected(onPeriphConnected),
		gatt.PeripheralDisconnected(onPeriphDisconnected),
	)

	device.Init(onStateChanged)
}

func onStateChanged(device gatt.Device, state gatt.State) {
	switch state {
	case gatt.StatePoweredOn:
		fmt.Println("Scanning...")
		device.Scan([]gatt.UUID{}, false)
		return
	default:
		device.StopScanning()
	}
}

func getQuaternion(data string) (quaternion string) {
	// Implement this function based on the pywitmotion library's get_quaternion function.
	return
}

func handleData(data string) {
	msgs := strings.Split(data, "U")
	for _, msg := range msgs {
		q := getQuaternion(msg)
		if q != "" {
			msgsNum++
			fmt.Println(q)
		}
		if msgsNum >= 100 {
			break
		}
	}
}
