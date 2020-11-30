package main

import (
	"fmt"
	"time"
)

func main() {

	runtimeQueue := queue{}
	runtimeQueue.Init()
	go func() {
		for i := 0; i < 100000; i++ {

			// 实际上这是单条协程在跑着追加
			runtimeQueue.Submit(i)
		}
	}()
	runtimeQueue.Run()
}

const batchSize = 40

type queue struct {
	queue chan int
}

func (receiver *queue) Init() {
	receiver.queue = make(chan int, 2048)
}

func (receiver *queue) Submit(i int) {
	receiver.queue <- i
}

func (receiver *queue) Run() {

	// 只有一条协程再这个 Run 中循环
	// 这里不会出现并发的现象
	batchData := make([]int, 0)
	for {
		select {
		case msg := <-receiver.queue:

			batchData = append(batchData, msg)
			if len(batchData) != batchSize {
				break
			}

			receiver.batchSend(batchData)

			batchData = make([]int, 0)
		}
	}
}

func (receiver *queue) batchSend(batchData []int) {

	for _, i := range batchData {
		time.Sleep(10 * time.Millisecond)
		fmt.Print(i, " ")
	}
	fmt.Printf("\n")
}
