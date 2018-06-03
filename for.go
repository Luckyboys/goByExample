package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 8; j <= 10; j++ {
		fmt.Println("j = ", j)
	}

	k := 0
	for {
		if k++; k > 2 {
			break
		}

		fmt.Println("k = ", k)
	}

	odd()
}

func odd() {
	for value := 0; value <= 5; value++ {
		if value%2 == 0 {
			continue
		}

		fmt.Println("value = ", value)
	}
}
