package main

import (
	"fmt"
	"time"
)

func main() {

	useEqual()
	useLessAndEqual()
}

func useEqual() {
	startTime := time.Now()

	for i := 0; i < 100000000000; i++ {

		if i == 0 {
			continue
		}
	}

	endTime := time.Now()

	fmt.Printf("i == 0, used Time: %.2fs\n", endTime.Sub(startTime).Seconds())
}

func useLessAndEqual() {

	startTime := time.Now()

	for i := 0; i < 100000000000; i++ {

		if i <= 0 {
			continue
		}
	}

	endTime := time.Now()

	fmt.Printf("i <= 0, used Time: %.2fs\n", endTime.Sub(startTime).Seconds())
}
