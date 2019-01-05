package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {

	var i uint64
	rand.Seed(time.Now().UnixNano())
	incrementValue := rand.Uint64()
	fmt.Printf("incrementValue: %d\n", incrementValue)
	atomic.AddUint64(&i, uint64(incrementValue%5+1))
	fmt.Println(i)
}
