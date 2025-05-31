package main

import (
	"fmt"
	"time"
)

func main() {
	// Select statements are used to listen on multiple Go channels.
	// The first channel that receives a message will be executed.
	jobQueue1 := make(chan int)
	jobQueue2 := make(chan int)

	// Simulate no job being sent to queues to trigger timeout
	select {
	case msg := <-jobQueue1:
		fmt.Println("Received job for processing from queue1:", msg)
	case msg := <-jobQueue2:
		fmt.Println("Received job for processing from queue2:", msg)
	case <-time.After(2 * time.Second): // Timeout case
		fmt.Println("No message received. Exiting.")
		// The default case would make select non-blocking if uncommented.
		// default:
	}
}
