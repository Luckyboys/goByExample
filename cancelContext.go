package main

import (
	"context"
	"fmt"
)

func main() {

	parentContext, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()
	childContext, childCancel := context.WithCancel(parentContext)
	if childContext.Err() != nil {
		fmt.Println("child error: ", childContext.Err().Error())
		return
	}

	if parentContext.Err() != nil {
		fmt.Println("parent error: ", parentContext.Err().Error())
		return
	}

	childCancel()
	if childContext.Err() != context.Canceled {
		fmt.Println("child error: ", childContext.Err())
		return
	}

	if parentContext.Err() != context.Canceled {
		fmt.Println("parent error: ", parentContext.Err())
		return
	}
}
