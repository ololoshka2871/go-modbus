package main

import (
	"flag"
	"fmt"
	"github.com/dpapathanasiou/go-modbus"
	"log"
)

func main() {

	// get the device serial port from the command line
	var (
		serialPort  string
		slaveDevice int
		startAddr   int
		numBytes    int
	)

	const (
		defaultPort         = ""
		defaultSlave        = 1
		defaultStartAddress = 3030
		defaultNumBytes     = 16
	)

	flag.StringVar(&serialPort, "serial", defaultPort, "Serial port (RS485) to use, e.g., /dev/ttyS0 (try \"dmesg | grep tty\" to find)")
	flag.IntVar(&slaveDevice, "slave", defaultSlave, fmt.Sprintf("Slave device number (default is %d)", defaultSlave))
	flag.IntVar(&startAddr, "start", defaultStartAddress, fmt.Sprintf("Start address (default is %d)", defaultStartAddress))
	flag.IntVar(&numBytes, "bytes", defaultNumBytes, fmt.Sprintf("Number of bytes to read from the start address (default is %d)", defaultNumBytes))
	flag.Parse()

	if len(serialPort) > 0 {

		// turn on the debug trace option, to see what is being transmitted
		trace := true

		// attempt to read the [startAddr] address register on
		// slave device number [slaveDevice] via the [serialDevice]
		readResult, readErr := modbusclient.RTURead(serialPort, byte(slaveDevice), modbusclient.FUNCTION_READ_HOLDING_REGISTERS, uint16(startAddr), uint16(numBytes), trace)
		if readErr != nil {
			log.Println(readErr)
		}
		log.Println(readResult)

	} else {

		// display the command line usage requirements
		flag.PrintDefaults()

	}

}
