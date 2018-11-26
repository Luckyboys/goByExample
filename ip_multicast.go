package main

import (
	"fmt"
	"net"
)

func main() {

	ip := "118.31.14.255"
	ipObject := net.ParseIP(ip)
	if ipObject == nil {

		fmt.Println("ipObject nil")
		return
	}

	if
	ipObject.IsUnspecified() ||
		ipObject.IsLoopback() ||
		ipObject.IsMulticast() ||
		ipObject.IsLinkLocalUnicast() ||
		ipObject.IsLinkLocalMulticast() {

		fmt.Println("in conditions")
	}

	fmt.Println("validate over")
}
