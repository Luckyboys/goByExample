package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"time"
)

func main() {

	duration, err := time.ParseDuration("1h")
	if err != nil {

		fmt.Println(err)
		return
	}

	data := struct {
		Duration time.Duration
	}{
		Duration: duration,
	}
	byteData, err := json.Marshal(data)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(byteData))
}
