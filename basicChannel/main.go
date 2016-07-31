package main

import "fmt"

//Actor model, Actor == the channel
//the channel ensure at a time, there will only be either side of the channel can access the channel
func main() {
	// deadlock()
	// deadlock2()
	bufferedChannel()
}

func deadlock() {
	ch := make(chan string)
	//waiting for a msg forever
	fmt.Println(<-ch)
}

func deadlock2() {
	ch := make(chan string)
	//since this channel has no capacity (non-buffered channel), this is blocking until someone receives the msg
	ch <- "Hello"
	fmt.Println(<-ch)
}

func bufferedChannel() {
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
}
