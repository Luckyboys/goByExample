package main

import (
	"time"
	"fmt"
	"github.com/Luckyboys/goByExample/lib"
)

func main() {

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	time1 := time.Now()

	time2 := time.Date(2018, time.March, 22, 19, 0, 0, 0, location)

	fmt.Println(time1)
	fmt.Println(time2)

	fmt.Printf("%#v\n", lib.DateDiff(
		&time1,
		&time2,
	))
}
