package main

import "fmt"

type animal interface {
	walk()
	fly()
}

type Duck struct {
}

type Chicken struct {
}

func (d Duck) walk() {
	fmt.Println("I'am a duck , ya ya")
}

func (d Duck) fly() {
	fmt.Println("I'm flying~~~")
}

func (c Chicken) walk() {
	fmt.Println("I'am a chicken , bo bo bo")
}

func (c Chicken) fly() {
	fmt.Println("I can't fly, I'm a chicken")
}

func letItFly(a animal) {
	a.fly()
}

func letItWalking(a animal) {
	a.walk()
}

func main() {

	d := Duck{}
	c := Chicken{}

	letItFly(d)
	letItFly(c)
	letItWalking(c)
	letItWalking(d)
}
