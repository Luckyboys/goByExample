package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println("2^10 = ", power(2, 10))
}

func power(d, s int) int {
	return int(math.Pow(float64(d), float64(s)))
}
