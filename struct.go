package main

import "fmt"

type User struct {
	userId   int
	username string
}

func main() {

	fmt.Println(User{100001, "Luckyboys"})
	fmt.Println(User{userId: 100002})
	fmt.Println(User{userId: 100003, username: "Xiao Kuang"})
	fmt.Println(&User{userId: 100004, username: "Jing Hao"})

	user := User{userId: 100005, username: "Li Lei"}
	fmt.Println(user.username)

	userPointer := &user
	fmt.Println(userPointer.userId)

	userPointer.username = "Hai Yang"
	fmt.Println(user.username)
}
