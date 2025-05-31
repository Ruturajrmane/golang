package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go goRoutines(ch)

	for val := range ch {
		fmt.Println("Recieved the message", val)
	}

}

func goRoutines(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Sending the into channel", i)
		ch <- i
		//On introdcung the sleep as the message is send it will be read by main goroutine immedately
		//Without sleep goroutine will write whole message latter this will be read by main thread
		//in the iterating
		time.Sleep(time.Second)
	}
	close(ch)
}
