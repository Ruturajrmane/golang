package main

import (
	"fmt"
	"time"
)

func main() {
	//buffered channel, the size of the buffer is mentioned in make
	//Unlike unbuffered channels, a buffered channel allows the sender
	// to send a message without blocking, as long as the buffer is not full.
	jobs := make(chan int, 5)
	done := make(chan bool)

	go processJob(jobs, done)

	for i := 1; i <= 5; i++ {
		fmt.Println("Sending the job for processing", i)
		jobs <- i
	}
	//closing is required to tell that all jobs are send
	close(jobs)
	//below line will block/wait the main thread to complete process jobs
	<-done
}

func processJob(ch chan int, done chan bool) {
	for {
		//more will become false when ch is closed & the all messages are drained
		job, more := <-ch
		if more {
			fmt.Println("Processing the job", job)
			time.Sleep(time.Second)
		} else {
			fmt.Println("All Jobs are done processigng")
			done <- true
			break
		}
	}
}
