package main

import (
	"fmt"
)

type Structure struct {
	AString string
}

func main() {

	var a *Structure
	s := new(Structure)
	a = s
	s.AString = "bbbb"

	fmt.Printf("%#v\n", a)
}
