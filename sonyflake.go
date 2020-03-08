package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	idCreator := sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (u uint16, err error) {
		// return uint16(rand.Int() % 65536), nil
		return 1, nil
	}})

	id, err := idCreator.NextID()
	if err != nil {
		logrus.WithError(err).Error("create id error")
		return
	}

	idInformation := sonyflake.Decompose(id)
	offsetTime, err := time.ParseDuration(fmt.Sprintf("%ds", int64(idInformation["time"]/100)))
	if err != nil {
		logrus.WithError(err).Error("offsetTime parse error")
		return
	}
	createTime := time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC).Add(offsetTime)
	fmt.Printf("id: %d , createTime: %s\n", id, createTime.Format("2006-01-02 15:04:05"))
}
