package main

import (
	"fmt"
	"sort"
)

func main() {

	aArray := []string{"a", "b"}
	bArray := []string{"b"}
	cArray := []string{"b", "a"}

	fmt.Println(`[]string{"a","b"} == []string{"a", "b"} ? `, fmt.Sprintf("%s", isEqualArray(aArray, aArray)))
	fmt.Println(`[]string{"a","b"} == []string{"b"} ? `, fmt.Sprintf("%s", isEqualArray(aArray, bArray)))
	fmt.Println(`[]string{"a","b"} == []string{"b", "a"} ? `, fmt.Sprintf("%s", isEqualArray(aArray, cArray)))
}

func isEqualArray(aArray []string, bArray []string) bool {

	length := len(aArray)
	if length != len(bArray) {
		return false
	}

	sort.Strings(aArray)
	sort.Strings(bArray)

	for i := 0; i < length; i++ {

		if aArray[i] != bArray[i] {
			return false
		}
	}

	return true
}
