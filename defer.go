package main

import (
	"fmt"
)

func main(){

	funcA()
}

func funcA(){

	defer func(){
		fmt.Printf("first\n")
	}()
	defer func(){
		fmt.Printf("second\n")
	}()
	defer func(){
		fmt.Printf("third\n")
	}()
}