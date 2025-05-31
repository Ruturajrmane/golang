package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type clients struct {
	lastLeakedTime time.Time
	tokens         int
}

var (
	capacity  = 10               //burst
	leakyRate = 12 * time.Second //5 request per minutes
	clientMap = make(map[string]*clients)
	mu        sync.Mutex
)

func middlewareLeakyRatelimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.ClientIP()
		mu.Lock()
		client, ok := clientMap[host]
		if !ok {
			newClient := &clients{
				lastLeakedTime: time.Now(),
				tokens:         1,
			}
			clientMap[host] = newClient
			fmt.Printf("New client added for the host : %v\n", host)
		} else {
			elapsedTime := time.Since(client.lastLeakedTime)
			leakedTokens := int(elapsedTime / leakyRate)

			if leakedTokens > 0 {
				client.tokens -= leakedTokens
				if client.tokens < 0 {
					client.tokens = 0
					fmt.Printf("Updated tokesn present for the host %v to %v\n", host, 0)
				}
				fmt.Printf("Updated tokesn present for the host %v to %v\n", host, client.tokens)
				client.lastLeakedTime = time.Now()
			}
			client.tokens++
			if client.tokens > capacity {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "Exceeded the rate limiter, Please try latter"})
				c.Abort()
				return
			}
			fmt.Printf("Updated tokesn present for the host %v to %v\n", host, client.tokens)

		}
		mu.Unlock()
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(middlewareLeakyRatelimiter())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":9090")
}
