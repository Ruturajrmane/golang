package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "JobId", "143")
	processJob(ctx)

}

func processJob(ctx context.Context) {
	value := ctx.Value("JobId")
	fmt.Println("Processed the job having id: ", value)
}
