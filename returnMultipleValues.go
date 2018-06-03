package main

import "fmt"

func main() {

	sum, multiple := compute(2, 4)
	fmt.Println("2 + 4 =", sum)
	fmt.Println("2 * 4 =", multiple)
}

func compute(digits ...  int) (int, int) {

	var sum, multiple int
	multiple = 1
	for _, digit := range digits {
		sum += digit
		multiple *= digit
	}

	return sum, multiple
}
