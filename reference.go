package main

import (
	"fmt"
)

type ss struct {
	str string
}

func main() {

	A := ss{str: "A"}
	B := ss{str: "B"}
	sList := []ss{A, B}
	rList := make([]*ss, 0)
	for _, item := range sList {
		newItem := item
		rList = append(rList, &newItem)
	}

	for _, item := range rList {
		fmt.Println(*item)
	}
}
