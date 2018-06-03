package main

import (
	"fmt"
	"math"
)

const s = "Ahahaha"

func main() {
	fmt.Println(s)
	const n = 1234
	const d = n / 33
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
