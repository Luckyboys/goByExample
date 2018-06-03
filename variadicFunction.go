package main

import (
	"fmt"
	strings2 "strings"
)

func main() {

	fmt.Println("1 + 1 =", sum(1, 1))
	fmt.Println("1 + 2 + 3 =", sum(1, 2, 3))
	fmt.Printf("%s = %d\n",func() (string) {
		var strings = make([]string, 10)
		var notFirst bool
		for i := 0; i < 10; i++ {
			if notFirst {
				strings = append(strings, " + ")
			}
			strings = append(strings, fmt.Sprint(i+1))
			notFirst = true
		}
		return strings2.Join(strings, "")
	}(), sum(func() (returnValue []int) {
		returnValue = make([]int, 10)
		for i := 1; i <= 10; i++ {
			returnValue[i-1] = i
		}
		return
	}()...))
}

func sum(digits ... int) int {
	var sum int
	for _, digit := range digits {
		sum += digit
	}
	return sum
}
