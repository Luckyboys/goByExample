package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	sig := <-chanSignal
	fmt.Printf("Received signal: %d\n", sig)

	fmt.Println("loop exited")
	return
}
