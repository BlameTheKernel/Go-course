package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort2(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Select an option:")
	fmt.Println("1 - Scan an specific port")
	fmt.Println("2 - Scan all ports")
	fmt.Println("3 - Scan a range of ports")

	reader.Scan()
	optionString := reader.Text()
	fmt.Println("You selected option ", optionString)

	flag.Parse()
	//optionInt := flag.Arg(optionString)
	// string to int
	OptionNumber, err := strconv.Atoi(optionString)
	fmt.Println(OptionNumber)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	switch OptionNumber {
	case 1: // name
		fmt.Println("Port Scanning")
		open := scanSpecificPort()
		fmt.Printf("Port Open: %t\n", open)

	case 2:
		fmt.Println("Scanning all ports")
		open := scanAllPorts()
		fmt.Printf("Port Open: %t\n", open)

	case 3:
		fmt.Println("Scanning all ports")
		open := scanRangePorts()
		fmt.Printf("Port Open: %t\n", open)

	default:
		fmt.Printf("Wrong parametes passed")
	}

	/*flag.Parse()
		portString := flag.Arg(0)
		// string to int
		portNumber, err := strconv.Atoi(portString)
		fmt.Println(portString, portNumber)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}


	//	flag.IntVar(&port, "p", 0, "Specify port")

	//	fmt.Println(port)

		fmt.Println("Port Scanning")
		open := scanPort("tcp", "localhost", 1313)
		fmt.Printf("Port Open: %t\n", open)*/
}
