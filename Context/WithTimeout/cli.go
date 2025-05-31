package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//this cancel can also be used to manually cancel suppose if the error occurs during the
	//processing & go routine exits even before time then, it's usedfull to release resource
	defer cancel()

	wg.Add(1)
	go processJob(ctx, &wg)
	wg.Wait()
}

func processJob(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		//After 5 seconds sleep from default in next iteration it will read this signal
		//and exit from the function
		case <-ctx.Done():
			fmt.Println("Exiting the process job due to context cancellation")
			return
		default:
			fmt.Println("Processing job")
			time.Sleep(5 * time.Second)
		}
	}
}
