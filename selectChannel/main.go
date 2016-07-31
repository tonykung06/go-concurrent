package main

import "fmt"

func main() {
	// basic()
	// deadlock()
	// eitherOne()
	// randomCaseWin()
	defaultCase()
}

func basic() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	msg := Message{
		To:      []string{"tony@tony.com"},
		From:    "tony2@tony.com",
		Content: "sample message",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message error out",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	fmt.Println(<-msgCh)
	fmt.Println(<-errCh)
}

func deadlock() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	}
}

func eitherOne() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"tony@tony.com"},
		From:    "tony2@tony.com",
		Content: "sample message",
	}

	msgCh <- msg

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	}
}

func randomCaseWin() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	msg := Message{
		To:      []string{"tony@tony.com"},
		From:    "tony2@tony.com",
		Content: "sample message",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message error out",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	}
}

func defaultCase() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No msg or error")
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
