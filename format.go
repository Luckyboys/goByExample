package main

import (
	"fmt"
	"time"
)

func main() {

	b := []byte("test")
	fmt.Printf("%%b: %b\n", b)
	fmt.Printf("%%s: %s\n", b)
	fmt.Printf("string(b): %v\n", string(b))

	fmt.Printf("%s\n", time.Now().Format(time.RFC3339))
}
