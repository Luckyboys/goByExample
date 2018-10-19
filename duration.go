package main

import (
	"fmt"
	"time"
)

func main() {

	var duration time.Duration
	duration = time.Duration(60020000000000)
	fmt.Printf("%.0fs\n", duration.Seconds())

	duration, err := time.ParseDuration(fmt.Sprintf("%.0fs", duration.Seconds()))
	fmt.Println(duration)
	fmt.Println(err)
}
