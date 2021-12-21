package main

import (
	"encoding/json"
	"fmt"
)

type Interface interface {
	String() string
}

type Data struct {
	AppId int64  `json:"appId"`
	Name  string `json:"name"`
}

func (d *Data) String() string {
	return fmt.Sprintf("appId: %d, name: %s", d.AppId, d.Name)
}

func main() {

	jsonString := `{"appId": 1, "name": "test"}`
	data := Data{}
	decodeData(jsonString, &data)

	fmt.Printf("data: %#v\n", data)
}

func decodeData(jsonString string, d Interface) {

	err := json.Unmarshal([]byte(jsonString), &d)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
