package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.DebugLevel)

	mutex := sync.RWMutex{}
	go func() {
		mutex.RLock()
		logrus.Debug("read locked 1")
		defer func() {
			mutex.RUnlock()
			logrus.Debug("read unlocked 1")
		}()
		time.Sleep(5 * time.Second)
	}()
	go func() {
		mutex.RLock()
		logrus.Debug("read locked 2")
		defer func() {
			mutex.RUnlock()
			logrus.Debug("read unlocked 2")
		}()
		time.Sleep(5 * time.Second)
	}()
	go func() {
		time.Sleep(time.Second)
		mutex.Lock()
		logrus.Debug("write locked")
		defer func() {
			mutex.Unlock()
			logrus.Debug("write unlocked")
		}()
		time.Sleep(time.Second)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		mutex.RLock()
		logrus.Debug("read locked 3")
		defer func() {
			mutex.RUnlock()
			logrus.Debug("read unlocked 3")
		}()
	}()

	time.Sleep(8 * time.Second)
}
