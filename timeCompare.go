package main

import (
	"time"
	"log"
)

const TimeFormat = "2006-01-02 15:04:05"

func main() {

	t1, err := time.Parse(TimeFormat, "2018-06-01 00:00:00")
	if err != nil {
		log.Println(err)
	}

	t2, err := time.Parse(TimeFormat, "2018-06-01 00:00:00")
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v.After( %v ) : %v", t1, t2, t1.After(t2))
	log.Printf("%v.Before( %v ) : %v", t1, t2, t1.Before(t2))
}
