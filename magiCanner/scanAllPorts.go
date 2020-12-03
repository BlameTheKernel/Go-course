package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanAllPorts() bool {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter IP to scan")
	reader.Scan()
	IP := reader.Text()

	fmt.Println("Choose protocol:")
	fmt.Println("1 - TCP")
	fmt.Println("2 - UDP")
	reader.Scan()
	protocol := reader.Text()
	protocolOption, err := strconv.Atoi(protocol)

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	switch protocolOption {
	case 1: // name
		protocol = "tcp"

	case 2:
		protocol = "udp"
	}
	fmt.Println("Scanning ports 1 to 65365 in ", protocol)

	for i := 1; i <= 1000; i++ {
		address := IP + ":" + strconv.Itoa(i)
		conn, err := net.DialTimeout(protocol, address, 60*time.Second)

		if err != nil {
			continue

		} else {
			fmt.Println("Port ", i, "opened")
		}

		defer conn.Close()
	}

	return true

}
