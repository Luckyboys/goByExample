package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)

	zeroValue(i)
	fmt.Println(i)

	zeroPointer(&i)
	fmt.Println(i)

	fmt.Println("pointer:", &i)
}

func zeroValue( i int ) {
	i = 0
}

func zeroPointer( p *int ) {
	*p = 0
}