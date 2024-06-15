This code defines a WT61Packet struct that represents a packet of data from the WT61 sensor. The ReadWT61Packet function reads a packet from an io.Reader (which could be a serial port, for example). The Data1 method converts the Data1H and Data1L fields to a signed 16-bit integer, and the Checksum method calculates the checksum of the packet. The IsValid method checks if the packet’s checksum matches the Check field.

Please replace getYourDataSource() with your actual data source. This is just a placeholder.

Remember to handle errors appropriately in your actual code. This is a simplified example and doesn’t include comprehensive error handling. Also, this code assumes that the data is sent in little-endian format. If your data is in a different format, you may need to adjust the code accordingly.


package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type WT61Packet struct {
	Header byte
	Type   byte
	Data1H byte
	Data1L byte
	Check  byte
}

func ReadWT61Packet(r io.Reader) (*WT61Packet, error) {
	var packet WT61Packet
	err := binary.Read(r, binary.LittleEndian, &packet)
	if err != nil {
		return nil, err
	}
	return &packet, nil
}

func (p *WT61Packet) Data1() int16 {
	return int16(p.Data1H)<<8 | int16(p.Data1L)
}

func (p *WT61Packet) Checksum() byte {
	return p.Header + p.Type + p.Data1H + p.Data1L
}

func (p *WT61Packet) IsValid() bool {
	return p.Checksum() == p.Check
}

func main() {
	// Replace with your actual data source
	r := getYourDataSource()

	packet, err := ReadWT61Packet(r)
	if err != nil {
		fmt.Println("Error reading packet:", err)
		return
	}

	if !packet.IsValid() {
		fmt.Println("Invalid packet")
		return
	}

	fmt.Println("Data1:", packet.Data1())
}



