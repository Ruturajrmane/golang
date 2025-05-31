package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var numberOfRequest = 5      //Number of request
var timePeriod = time.Minute //5 request per minute are allowed
var rateLimiterMap = make(map[string]*rateLimiter)
var mu sync.Mutex

type rateLimiter struct {
	requests    int
	startWindow time.Time
}

func middlewareRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		fmt.Printf("Made a request from ip : %v\n", ip)

		mu.Lock()

		client, ok := rateLimiterMap[ip]

		if !ok || time.Since(client.startWindow) > timePeriod {
			newClient := &rateLimiter{
				requests:    1,
				startWindow: time.Now(),
			}
			fmt.Println("Creating a new client")
			rateLimiterMap[ip] = newClient
		} else {
			if client.requests >= numberOfRequest {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "Number of requests exeeeced, Please try letter after some time"})
				c.Abort() //can be combined with AbortWithJson
				return
			} else {
				client.requests++
				fmt.Printf("Totatl %d requests are hit\n", client.requests)
			}
		}
		mu.Unlock()
		c.Next()

	}
}

func main() {
	router := gin.Default()

	router.Use(middlewareRateLimiter())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	fmt.Println("Starting a new server at 9090")
	http.ListenAndServe(":9090", router)
}
