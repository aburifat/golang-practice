package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		time.Sleep((2 * time.Second) - (1 * time.Millisecond))
		c <- "Hello, World!"
		close(c)
	}()

	select {
	case res := <-c:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}
