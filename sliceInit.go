package main

import (
	"fmt"
)

type testStruct struct {
}

func main() {

	// panic: runtime error: index out of range [0] with length 0
	// var list = make([]*testStruct, 0, 10)
	// list[0] = &testStruct{}

	// it will work
	var list = make([]*testStruct, 10)
	list[0] = &testStruct{}
	for index, item := range list {

		if item == nil {
			fmt.Printf("index: %d is nil\n", index)
		} else {
			fmt.Printf("index: %d had value\n", index)
		}
	}
}
