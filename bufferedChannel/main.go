package main

import (
	"fmt"
	"strings"
)

func main() {
	// bufferedChannel()
	// closingChannel()
	// loopingOverChannel()
	rangingOverChannel()
}

func bufferedChannel() {
	phrase := "Hey, how is a buffered channel working?"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}

func closingChannel() {
	phrase := "Hey, how is a buffered channel working?"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	//close the sending of the channel, the data already inside the channel will still be there
	//close the channel to tell the receiver that no more data
	close(ch)
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	//panic: send on closed channel
	// ch <- "test"
}

func loopingOverChannel() {
	phrase := "Hey, how is a buffered channel working?"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	//close the channel so that ok will be false
	close(ch)
	for {
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
}

func rangingOverChannel() {
	phrase := "Hey, how is a buffered channel working?"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	//close the channel so that ok will be false
	close(ch)

	for msg := range ch {
		fmt.Print(msg + " ")
	}
}
