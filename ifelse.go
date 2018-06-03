package main

import "fmt"

func main() {
	if num := 10; num%3 == 1 {
		fmt.Println("1: ", num)
	} else if num%3 == 2 {
		fmt.Println("2: ", num)
	} else {
		fmt.Println("0: ", num)
	}
	//compile error 作用域只有在if里面
	//fmt.Println( "m: " , num )
}
