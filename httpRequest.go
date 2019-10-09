package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	body := "snapshot"

	req, err := http.NewRequest("PUT", "http://www.kpaas.net", strings.NewReader(body))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	req.Body, err = req.GetBody()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("body: %s\n", string(data))
	_ = req.Body.Close()

	req.Body, err = req.GetBody()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	data, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("body: %s\n", string(data))
}
