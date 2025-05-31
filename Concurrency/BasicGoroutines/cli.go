package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting the main thread")
	var wg sync.WaitGroup //This is a empty wg not initialized(No memory allocation) pointer of this will point to nil
	wg.Add(1)             //Initialization is done here
	go hello(&wg)
	wg.Wait()

	fmt.Println("Hello from the thread")
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from the Go routine")
}
