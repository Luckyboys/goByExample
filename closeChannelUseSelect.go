package main

import (
	"fmt"
	"sync"
)

var channelSelect chan int
var wgSelect sync.WaitGroup

func main() {

	channelSelect = make(chan int)
	go producer()
	for i := 0; i < 5; i++ {
		wgSelect.Add(1)
		go consumer()
	}

	wgSelect.Wait()
}

func producer() {

	for i := 0; i < 10; i++ {

		channelSelect <- i
	}

	close(channelSelect)
}

func consumer() {

loop:
	for {
		select {
		case item, ok := <-channelSelect:

			if !ok {
				fmt.Println("completed")
				break loop
			}

			fmt.Println(item)
		}
	}

	wgSelect.Done()
}
