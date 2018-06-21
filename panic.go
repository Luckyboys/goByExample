package main

import (
	"log"
)

func main() {

	err := recover()
	log.Println("recover", err)

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("ahahaha don't panic")
}