package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("%d\n", uint64(0xffffffffffffffff))
	fmt.Printf("%s\n", strconv.FormatUint(uint64(0xffffffffffffffff), 10))
}
