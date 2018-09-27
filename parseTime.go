package main

import (
	"github.com/Luckyboys/goByExample/lib"
	"log"
	"time"
)

func main() {

	var timeStamp int64
	var ok bool
	timeStamp, ok = lib.ParseTime(time.Unix(1538039361, 0))
	log.Printf("UnixTimeStamp: %d , ParseOK: %t\n", timeStamp, ok)

	timeStamp, ok = lib.ParseTime(1538039361)
	log.Printf("UnixTimeStamp: %d , ParseOK: %t\n", timeStamp, ok)

	timeStamp, ok = lib.ParseTime("2018-12-08 23:59:59")
	log.Printf("UnixTimeStamp: %d , ParseOK: %t\n", timeStamp, ok)

	timeStamp, ok = lib.ParseTime("2018-12-08T15:59:59Z")
	log.Printf("UnixTimeStamp: %d , ParseOK: %t\n", timeStamp, ok)

	timeStamp, ok = lib.ParseTime("20181208235959")
	log.Printf("UnixTimeStamp: %d , ParseOK: %t\n", timeStamp, ok)
}
