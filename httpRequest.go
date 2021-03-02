package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	getTwiceBody()

	sendWithTimeout()

	httpRequestURI()
}

func httpRequestURI() {

	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com/?c=2&a=2&b=1", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("requestURI: %s\n", req.URL.RequestURI())
}

func getTwiceBody() {
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

func sendWithTimeout() {

	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Printf("start request\n")
	response, err := http.DefaultClient.Do(req.WithContext(ctx))
	fmt.Printf("request compeleted\n")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	defer func() {
		_ = response.Body.Close()
	}()

	fmt.Printf("status: %d\n", response.StatusCode)
}
