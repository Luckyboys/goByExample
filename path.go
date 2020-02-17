package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	fmt.Printf("path.Clean(\"./\") = %s\n", path.Clean("./"))
	dir, err := os.Getwd()
	fmt.Printf("os.Getwd() = %s, error = %v\n", dir, err)
}
