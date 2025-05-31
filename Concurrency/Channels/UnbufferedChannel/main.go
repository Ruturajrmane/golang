package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel.
	// Unbuffered channels are synchronous: a send operation will block
	// until another goroutine is ready to receive, and vice versa.
	ch := make(chan string)

	// Start a goroutine that will send a message to the channel
	go greeting(ch)

	// Receive from the channel. This blocks until the goroutine sends the message.
	fmt.Println(<-ch)

	// This will execute after the message is received.
	fmt.Println("Good Morning from main go routine")

}

func greeting(ch chan string) {
	// Send a message to the channel. This blocks until the main goroutine receives it.
	ch <- "Good Morning from go routine"
}
