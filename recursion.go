package main

import "fmt"

func main() {

	fmt.Println("20! =", fact(uint64(20)))
}

func fact(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
