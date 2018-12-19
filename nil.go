package main

import (
	"fmt"
	"github.com/qbox/qvm/server/errors"
)

func main() {

	var err *errors.Error
	fmt.Printf("error: %#v\n", err)
	err = nil
	fmt.Printf("error: %#v\n", err)

	if err == nil {

		fmt.Printf("%#v == nil\n", err)
	} else {

		fmt.Printf("%#v != nil\n", err)
	}
}
