package main

import (
	"fmt"
)

func main() {

	var err error
	fmt.Printf("error: %#v\n", err)
	err = nil
	fmt.Printf("error: %#v\n", err)

	if err == nil {

		fmt.Printf("%#v == nil\n", err)
	} else {

		fmt.Printf("%#v != nil\n", err)
	}
}
