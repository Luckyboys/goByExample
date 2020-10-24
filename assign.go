package main

import (
	"fmt"
)

var storedValue interface{}

func getValueByRef() (value interface{}) {

	return storedValue
}

func set(value interface{}) {
	storedValue = value
}

func main() {

	var str string
	set("abc")
	str = getValueByRef().(string)
	fmt.Println(str)
}
