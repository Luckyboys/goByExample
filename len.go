package main

import (
	"fmt"
)

func main() {

	fmt.Printf("len(\"中文\"): %d\n", len("中文"))
	fmt.Printf("len([]rune(\"中文\")): %d\n", len([]rune("中文")))
	fmt.Printf("len(\"english\"): %d\n", len("english"))
	fmt.Printf("len([]rune(\"中文+english\")): %d\n", len([]rune("中文+english")))
}
