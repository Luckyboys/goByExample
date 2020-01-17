package main

import (
	"fmt"
)

func main() {

	b := []byte("test")
	fmt.Printf("%%b: %b\n", b)
	fmt.Printf("%%s: %s\n", b)
	fmt.Printf("string(b): %v\n", string(b))
}
