package main

import (
	"fmt"
	"net"
)

func main() {

	address := "http://127.0.0.1:2379"

	splitAddress, splitPort, splitErr := net.SplitHostPort(address)
	fmt.Printf("address: %s, port: %d, err: %v\n", splitAddress, splitPort, splitErr)

	if ae, ok := splitErr.(*net.AddrError); ok && ae.Err == "missing port in address" {

		splitPort = "2379"
		splitAddress = address
		fmt.Printf("set value: %s\n", net.JoinHostPort(splitAddress, splitPort))
	}
}
