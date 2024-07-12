[Wit CPP](https://github.com/WITMOTION/WitStandardProtocol_JY901/tree/main/C%2B%2B/NORMAL_WIN_CPP)

[C++ to GO](https://www.codeconvert.ai/c++-to-golang-converter)



package serial

import (
	"github.com/jacobsa/go-serial/serial"
	"io"
)

type Port struct {
	io.ReadWriteCloser
}

func Open(port string, baud int) (*Port, error) {
	options := serial.OpenOptions{
		PortName:        port,
		BaudRate:        uint(baud),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		return nil, err
	}

	return &Port{port}, nil
}
