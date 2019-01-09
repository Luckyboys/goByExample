package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	node, err := snowflake.NewNode(rand.Int63() % 1024)
	if err != nil {

		fmt.Println("Error: ", err)
		return
	}

	id := node.Generate()
	idInt64 := id.Int64()
	id = snowflake.ID(idInt64)
	createTime := time.Unix(id.Time()/1000, (id.Time()%1000)*1000)
	fmt.Printf("id: %d , createTime: %s\n", id.Int64(), createTime.Format("2006-01-02 15:04:05"))
}
