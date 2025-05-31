package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go processJob(ctx, &wg)

	time.Sleep(12 * time.Second)
	fmt.Println("Cancelling the job")
	cancel()
	wg.Wait()
}

func processJob(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Canceled job processing")
			return
		//below time out is unnecessary it will never trigger as with each iteration a new timer of 10 seconds
		//introduced if no select statement is executed then only it will be alive as there is
		//default it will get reset -> time.After is how much you wait without going inside any select or default
		case <-time.After(3 * time.Second):
			fmt.Println("Job processing timed out")
			return
		default:
			fmt.Println("Processing the job")
			//This is introduced to not run loop uncontrollablly or simulate real work
			time.Sleep(5 * time.Second)
		}
	}
}
