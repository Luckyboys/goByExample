package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/anaskhan96/go-password-encoder"
)

func main() {
	// Using the default options
	salt, encodedPwd := password.Encode("generic password", nil)
	check := password.Verify("generic password", salt, encodedPwd, nil)
	fmt.Println("password: ", encodedPwd)
	fmt.Println("salt: ", salt)
	fmt.Println(check) // true

	// Using custom options
	options := &password.Options{10, 10000, 50, sha1.New}
	salt, encodedPwd = password.Encode("generic password", options)
	check = password.Verify("generic password", salt, encodedPwd, options)
	fmt.Println("password: ", encodedPwd)
	fmt.Println("salt: ", salt)
	fmt.Println(check) // true
}
