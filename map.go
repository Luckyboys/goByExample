package main

import (
	"fmt"
	"time"
)

const (
	KeyUserId    = "userId"
	KeyLoginTime = "loginTime"
	KeySessionId = "sessionId"
)

func main() {

	user := make(map[string]interface{})

	fmt.Println(user)
	printMaps(user)

	user[KeyLoginTime] = "1234321"
	user[KeyUserId] = 123

	fmt.Println(user)
	printMaps(user)

	fmt.Println("map key count: ", len(user))

	userId := user[KeyUserId]
	fmt.Println("userId: ", userId)

	_, isSetSessionId := user[KeySessionId]
	fmt.Println("isSetSessionId: ", isSetSessionId)

	delete(user, KeyLoginTime)

	printMaps(user)
	fmt.Println("map key count: ", len(user))

	printMaps(map[string]interface{}{
		KeyUserId: "1392918432", KeySessionId: "d23e8328232432effaa", KeyLoginTime: time.Now().Unix()})
}

func printMaps(keyValuePairs map[string]interface{}) {

	fmt.Println("{")
	for key, value := range keyValuePairs {
		fmt.Printf("\t[%s]: %v\n", key, value)
	}
	fmt.Println("}")
}
