package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]int
	limit    int
	window   time.Duration
	mutex    sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
		mutex:    sync.Mutex{},
	}
}

func (rl *RateLimiter) resetCount(clientID string) {
	time.Sleep(rl.window)
	rl.mutex.Lock()
	delete(rl.requests, clientID)
	rl.mutex.Unlock()
}

func (rl *RateLimiter) AllowRequest(clientID string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Check the current count for the client
	count, exists := rl.requests[clientID]
	if !exists || count < rl.limit {
		// Reset the count for a new window or increment for the current window
		if !exists {
			go rl.resetCount(clientID)
		}
		rl.requests[clientID]++
		return true
	}

	return false
}

func main() {
	// Initialize the rate limiter: limit of 5 requests per 10 seconds
	rateLimiter := NewRateLimiter(5, 10*time.Second)

	// Simulate client requests
	clientID := "client1" // Example client ID

	for i := 0; i < 10; i++ {
		if rateLimiter.AllowRequest(clientID) {
			fmt.Printf("Request %d allowed for %s\n", i+1, clientID)
		} else {
			fmt.Printf("Request %d denied for %s\n", i+1, clientID)
		}

		// Wait a second between requests
		time.Sleep(1 * time.Second)
	}
}
