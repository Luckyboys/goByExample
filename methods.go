package main

import (
	"time"
	"fmt"
)

type Person struct {
	userId   int
	username string
	birthday int
}

func (user *Person) computeAge() int {

	return user.computeAge2(time.Now())
}

func (user Person) computeAge2(currentTime time.Time) int {

	birthday := time.Unix(int64(user.birthday), 0)

	age := currentTime.Year() - birthday.Year()

	switch {
	case birthday.Month() < currentTime.Month():
		age--
	case birthday.Month() == currentTime.Month() && birthday.Day() < currentTime.Day():
		age--
	}

	return age
}

func main() {

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	user := Person{userId: 100001, username: "Luckyboys", birthday: int(time.Date(1984, time.December, 8, 0, 0, 0, 0, location).Unix())}
	fmt.Println(user)
	fmt.Println(user.computeAge())

	userPointer := &user
	fmt.Println(userPointer)
	fmt.Println(userPointer.computeAge2(time.Now()))
}
