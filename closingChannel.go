package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			i, ok := <-jobs
			if ok {
				fmt.Println("received", i)
				time.Sleep( time.Second )
			} else {
				fmt.Println("overed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent", i)
	}

	close(jobs)
	fmt.Println("sent all")

	<-done
}
