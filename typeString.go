package main

import (
	"fmt"
)

type AType string
type BType string

const A1Type AType = "A1"
const A2Type AType = "A2"

const B1Type BType = "B1"
const B2Type BType = "B2"

type A struct {
	AType AType
}

type B struct {
	BType BType
}

func main() {

	a := &A{AType: A1Type}
	b := &B{BType: B1Type}

	if a.AType == b.BType {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
