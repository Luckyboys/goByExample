package main

import (
	"fmt"
	"io"
	"time"
)

var times int

func receivedData() *recvData {

	times++
	if times <= 10 {
		return &recvData{
			d: []byte{byte(times)},
		}
	}

	time.Sleep(10 * time.Second)
	return &recvData{
		e: io.EOF,
	}
}

type recvData struct {
	d []byte
	e error
}

func main() {

	closeChannel := make(chan bool, 1)
	var stopped bool
	stop := func() {
		stopped = true
		closeChannel <- true
		close(closeChannel)
	}

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("timeout")
		stop()
	}()
	dataChannel := make(chan *recvData)
	go func() {
		for {
			if stopped {
				break
			}
			dataChannel <- receivedData()
		}
	}()

loop:
	for {
		select {
		case d := <-dataChannel:
			fmt.Printf("%#v\n", d)
			if d.e == io.EOF {
				fmt.Println("EOF")
				stop()
				continue
			}

		case <-closeChannel:
			fmt.Println("break")
			break loop
		}
	}

	close(dataChannel)

	fmt.Println("closed application")
}
