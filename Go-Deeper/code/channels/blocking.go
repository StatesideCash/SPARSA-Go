package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Using channels to join on a goroutine
	ci := make(chan int)
	fmt.Printf("Pinging google... ")
	go func() {
		// Ping google and wait for completion
		exec.Command("/bin/ping", "-c 4", "8.8.8.8").Run()
		ci <- 1
	}()
	// Channel receiver will block until it gets data
	<-ci
	fmt.Println("Done!")
}
