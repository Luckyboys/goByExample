package main

import (
	"runtime"
	"fmt"
)

func main() {

	for i := 0 ; i < 10000000 ; i++ {
		c := make(chan int,1000)
		for times := 0 ; times < 500 ; times++ {
			c <- i
		}

		PrintMemory()
	}
}


var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}