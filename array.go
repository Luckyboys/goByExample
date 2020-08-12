package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 12
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println(len(a[:2]))
	fmt.Println(cap(a[:2]))

	b := a[:2]
	b[1] = 22

	fmt.Println(a)

	width := 10
	height := 5
	c := make([][]int, width)
	for x := 0; x < width; x++ {
		c[x] = make([]int, height)
		for y := 0; y < height; y++ {
			c[x][y] = x + y
		}
	}

	fmt.Println("[")
	for x := range c {
		fmt.Printf("\t%v\n", c[x])
	}
	fmt.Println("]")

	type S struct {
	}
	sList := make([]*S, 0)
	fmt.Println("empty array == nil ? %v", sList == nil)
}
