package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {

	a := bson.M{
		"resource_id": "32323-2222",
	}

	b := convert(a)

	b["soft_delete.was_deleted"] = false

	fmt.Println(bson.M(b))
}

func convert(obj map[string]interface{}) map[string]interface{} {

	return obj
}
