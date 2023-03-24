package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	fmt.Printf("case1: precall cancel: (error)%v\n", ctx.Err())
	cancel()
	fmt.Printf("case1: call cancel 1 time: (error)%v\n", ctx.Err())
	cancel()
	fmt.Printf("case1: call cancel 2 time: (error)%v\n", ctx.Err())
	time.Sleep(time.Second)
	fmt.Printf("case1: timeout: (error)%v\n", ctx.Err())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	time.Sleep(2 * time.Second)
	fmt.Printf("case2: timeout: (error)%v\n", ctx.Err())
	cancel()
	fmt.Printf("case2: call cancel after timeout: (error)%v\n", ctx.Err())
}
