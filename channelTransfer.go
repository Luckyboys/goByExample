package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	dataChannel := make(chan int,5)
	rwLock := sync.RWMutex{}
	isClosed := false

	for i := 0; i < 10; i++ {
		go func(i int) {
			rwLock.RLock()
			defer rwLock.RUnlock()

			if isClosed {
				return
			}

			dataChannel <- i
			fmt.Println(i, "input")
		}(i)
	}

	for i := 0; i < 5; i++ {

		func() {
			rwLock.RLock()
			defer rwLock.RUnlock()

			out, ok := <-dataChannel
			if !ok {
				fmt.Println("not ok")
				return
			}
			fmt.Println("out", out)
			time.Sleep(100 * time.Millisecond)
		}()
	}

	rwLock.Lock()
	close(dataChannel)
	rwLock.Unlock()

	time.Sleep(10 * time.Second)
}
