package main

import (
	"fmt"
)

var chanForEach chan int

func main() {

	chanForEach = make(chan int)
	go producerForEach()
	consumerForEach()
}

func producerForEach() {

	for i := 0; i < 10; i++ {

		chanForEach <- i
	}

	close(chanForEach)
}

func consumerForEach() {

	for item := range chanForEach {

		fmt.Println(item)
	}

}
