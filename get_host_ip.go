package main

import (
	"fmt"
	"net"
)

func main() {

	host := "qce-portal-poc.cloudappl.com"
	ips, err := net.LookupIP(host)
	if err != nil {
		return
	}

	if len(ips) <= 0 {
		fmt.Printf("resolve %s host error", host)
		return
	}

	for index, ip := range ips {
		fmt.Printf("ip[%d]: %#v\n", index, ip.String())
	}

}
