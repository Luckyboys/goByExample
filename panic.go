package main

import (
	"log"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("ahahaha don't panic")
}