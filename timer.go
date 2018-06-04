package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {
	timer := time.NewTimer(time.Second)

	for i := 0 ; i < 10000 ; i++ {
		go func() {
			//这是一个危险的用法。万一在外面stop了，而不是跑这里，go routine就不会结束，从而导致内存溢出
			<-timer.C
			fmt.Println("ahaha")
		}()
	}

	if timer.Stop() {
		fmt.Println("stopped")
	}

	PrintMemory()

	time.Sleep(time.Second * 2)
}

var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}