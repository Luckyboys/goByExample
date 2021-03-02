package main

import (
	"fmt"
	"net/url"
)

func main() {

	queryString := url.Values{}
	queryString.Set("appId", "28882828281")
	queryString.Set("键", "值")
	queryString.Set("special&character", "some%special%character")

	fmt.Printf("%s\n", queryString.Encode())
}
