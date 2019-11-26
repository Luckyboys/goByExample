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

	isSamePointer()
}

func zeroValue(i int) {
	i = 0
}

func zeroPointer(p *int) {
	*p = 0
}

type TestPointerStruct struct {
	Name string
}

func isSamePointer() {

	structA := new(TestPointerStruct)
	structB := new(TestPointerStruct)
	structCPointA := structA

	fmt.Printf("is same struct pointer structA == structB: %v\n", structA == structB)
	fmt.Printf("is same struct pointer structA == structCPointA: %v\n", structA == structCPointA)

	fmt.Printf("is same struct value: %v\n", *structA == *structB)
}
