package main

import (
	"fmt"
)

type Changed struct {
	isNewRecord bool `bson:"-" json:"-"`
	isUpdated   bool `bson:"-" json:"-"`
}

func (modelData *Changed) IsNewRecord() bool {
	return modelData.isNewRecord
}

func (modelData *Changed) Updated() bool {
	return modelData.isUpdated
}

func (modelData *Changed) SetUpdate(isUpdated bool) {
	modelData.isUpdated = isUpdated
}

func (modelData *Changed) SetNewRecord(isNewRecord bool) {
	modelData.isNewRecord = isNewRecord
}

type Master struct {
	UserId uint

	Changed
}

func main() {

	m := &Master{}

	fmt.Println(m.IsNewRecord())
	m.SetNewRecord(true)
	fmt.Println(m.IsNewRecord())
}
