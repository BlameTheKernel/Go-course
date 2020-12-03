package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanRangePorts() bool {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter initial port:")
	reader.Scan()
	firstPortStr := reader.Text()

	firstPort, err := strconv.Atoi(firstPortStr)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Enter last port:")
	reader.Scan()
	lastPortStr := reader.Text()

	lastPort, err := strconv.Atoi(lastPortStr)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

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
	fmt.Println("Scanning ports", firstPort, "to", lastPort, "in", protocol)

	for i := firstPort; i <= lastPort; i++ {
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
