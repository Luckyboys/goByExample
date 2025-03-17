package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	var parameters = make(map[string]any)
	// parameters["taskId"] = "3"
	taskId := cast.ToInt64(parameters["taskId"])
	if taskId <= 0 {
		fmt.Printf("taskId is zero\n")
	}
	
	fmt.Printf("taskId: %d\n", taskId)
}
