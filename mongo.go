package main

import (
	"github.com/Luckyboys/goByExample/lib"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var (
	databaseName   = "test"
	collectionName = "sort_table"
)

func main() {

	dsn := "mongodb://127.0.0.1:27017"
	session, err := mgo.Dial(dsn)
	if err != nil {

		log.Printf("Connect error %s\n", err)
		return
	}

	defer session.Close()

	c := session.DB(databaseName).C(collectionName)

	result := make(map[string]interface{})
	err = c.Find(bson.M{}).One(&result)

	if err != nil {
		log.Printf("Query error %s\n", err)
		return
	}

	log.Printf("%#v\n", result)
	log.Printf("UpdateTime: %s\n", result["update_time"])

	var timeStamp, ok = lib.ParseTime(result["update_time"])
	if ok {
		log.Printf("UnixTimeStamp: %d\n", timeStamp)
	} else {
		log.Printf("Parse failure\n")
	}
}
