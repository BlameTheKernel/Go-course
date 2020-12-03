package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanSpecificPort() bool {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter IP to scan")
	reader.Scan()
	IP := reader.Text()

	fmt.Println("Enter port to scan")
	reader.Scan()
	port := reader.Text()
	portNumber, err := strconv.Atoi(port)

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

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

	fmt.Println("Scanning IP:", IP, "\nPort:", portNumber, "\nProtocol:", protocol, "\n\n")
	address := IP + ":" + strconv.Itoa(portNumber)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		fmt.Println("Port ", portNumber, "closed")
		return false
	} else {
		fmt.Println("Port ", portNumber, "opened")
	}
	defer conn.Close()
	return true
}
