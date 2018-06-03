package main

import "fmt"

func main() {
	zero := 0
	fmt.Println(^zero)
	uZero := uint(zero)
	fmt.Println(^uZero)
	runeZero := rune(zero)
	fmt.Println(^runeZero)
}
