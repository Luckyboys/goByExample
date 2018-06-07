package main

import (
	"errors"
	"sync"
	"fmt"
	"math/rand"
	"go.uber.org/atomic"
	"time"
)

type Queue struct {
	data chan interface{}
	lock sync.Mutex
	isLocking *atomic.Bool
}

func NewQueue() *Queue {
	q := new(Queue)
	q.data = make(chan interface{},2000000)
	q.isLocking = atomic.NewBool(false)
	return q
}

func (q *Queue) Push(v interface{}) {

	q.extend()
	for q.isLocking.Load() {
		time.Sleep(time.Microsecond)
	}
	q.data <- v
}

func (q *Queue) Pop() (interface{}, error) {

	if len(q.data) > 0 {
		return <-q.data, nil
	}

	return nil, errors.New("no value in queue")
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Dump() []interface{} {

	queueLength := len(q.data)
	nextQueue := make(chan interface{}, queueLength)

	q.lock.Lock()
	q.isLocking.Store(true)
	defer q.lock.Unlock()
	defer q.isLocking.Store(false)

	var res []interface{}
	for i := 0 ; i < queueLength ; i ++ {
		ele := <- q.data
		nextQueue <- ele
		res = append(res, ele)
	}

	return res
}

func (q *Queue) Reset() {
	q.data = make(chan interface{})
}

func (q *Queue) extend() {

	if len(q.data) < cap(q.data) {
		return
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) < cap(q.data) {
		return
	}

	queueLength := len(q.data)
	newQueueLength := queueLength

	if queueLength == 0 {
		newQueueLength = 1
	}

	nextQueue := make(chan interface{}, newQueueLength*2)
	for i := 0 ; i < queueLength ; i++ {
		ele := <- q.data
		nextQueue <- ele
	}
	q.data = nextQueue
}

func main() {

	q := NewQueue()

	producerRoutine := atomic.NewUint64(0)
	consumerRoutine := atomic.NewUint64(0)

	for i := 0 ; i < 1000 ; i++ {

		producerRoutine.Add(1)
		go func(){
			for j := 0 ; j < 100000 ; j++ {
				q.Push(rand.Int())
			}
			producerRoutine.Sub(1)
		}()
	}

	for producerRoutine.Load() > 0 {

		fmt.Printf("Length: %d , Producer: %d , Consumer: %d\n",
			q.Len(), producerRoutine.Load(), consumerRoutine.Load(),
		)
		time.Sleep(time.Second)
	}

	for i := 0 ; i < 100 ; i++ {

		consumerRoutine.Add(1)
		go func(){
			for {
				_ , err := q.Pop()
				if err != nil {
					break
				}
			}
			consumerRoutine.Sub(1)
		}()
	}

	startTime := time.Now()
	for consumerRoutine.Load() > 0 || producerRoutine.Load() > 0 || time.Now().Sub(startTime).Minutes() < 1 {

		fmt.Printf("Length: %d , Producer: %d , Consumer: %d\n",
			q.Len(), producerRoutine.Load(), consumerRoutine.Load(),
		)
		time.Sleep(time.Second)
	}
}
