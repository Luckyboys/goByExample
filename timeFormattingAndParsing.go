package main

import (
	"time"
	"fmt"
)

func main() {

	t1 , err := time.Parse( "2006-01-02 15:04:05" , "2018-05-01 20:12:34" )


	if err != nil {
		panic( "time format error" )
	}
	fmt.Println(t1)
	fmt.Println(t1.Unix())
}