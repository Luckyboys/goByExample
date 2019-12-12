package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	testDurationJSON()
	testJSONMarshalError()
	testJSONUnmarshalNullString()
}

func testJSONMarshalError() {
	data, err := json.Marshal(struct {
		A interface{}
	}{make(chan int)})
	fmt.Printf("data: %v, err: %v\n", string(data), err)
}

func testDurationJSON() {
	duration, err := time.ParseDuration("1h")
	if err != nil {

		fmt.Println(err)
		return
	}
	data := struct {
		Duration time.Duration
	}{
		Duration: duration,
	}
	byteData, err := json.Marshal(data)
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(string(byteData))
}

func testJSONUnmarshalNullString() {

	jsonString := []byte(`{"data": null}`)
	data := struct {
		Data string `json:"data"`
	}{}
	err := json.Unmarshal(jsonString, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("data: %#v\n", data)
}
