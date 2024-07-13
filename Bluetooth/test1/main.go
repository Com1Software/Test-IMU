package main

import (
    "fmt"
    "net"
    "strings"
)

func getQuaternion(msg []byte) []float64 {
    return wit.GetQuaternion(msg)
}

// func getMagnetic(msg []byte) []float64 {
//     return wit.GetMagnetic(msg)
// }

// func getAngle(msg []byte) []float64 {
//     return wit.GetAngle(msg)
// }

// func getGyro(msg []byte) []float64 {
//     return wit.GetGyro(msg)
// }

// func getAcceleration(msg []byte) []float64 {
//     return wit.GetAcceleration(msg)
// }

func main() {
    // set your device's address
    imu := "00:0C:BF:02:1E:40"

    // Create the client socket
    socket, err := net.Dial("bluetooth", imu+":1")
    if err != nil {
        fmt.Println("Error connecting to device:", err)
        return
    }
    defer socket.Close()

    msgs_num := 0
    for msgs_num < 100 {
        data := make([]byte, 1024)
        n, err := socket.Read(data)
        if err != nil {
            fmt.Println("Error reading from socket:", err)
            return
        }

        // split the data into messages
        messages := strings.Split(string(data[:n]), "U")
        for _, msg := range messages {
            q := getQuaternion([]byte(msg))
            // q := getMagnetic([]byte(msg))
            // q := getAngle([]byte(msg))
            // q := getGyro([]byte(msg))
            // q := getAcceleration([]byte(msg))
            if q != nil {
                msgs_num++
                fmt.Println(q)
            }
        }
    }
}

