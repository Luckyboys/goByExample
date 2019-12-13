package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var err error

	err = ioutil.WriteFile("/tmp/test", []byte("test data"), os.FileMode(0644))
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println("Write completed")
}
