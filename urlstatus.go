package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel
	ch := make(chan string)

	// Launch a goroutine for each link to check its status
	for _, link := range links {
		go checkLink(link, ch)
	}

	// Continuously receive links from the channel and re-check them after a delay
	for l := range ch {
		// Launch a new goroutine to re-check the link after 3 seconds
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	start := time.Now() // Record start time for latency measurement

	resp, err := http.Get(link)
	duration := time.Since(start) // Calculate request duration

	if err != nil {
		fmt.Printf("%-30s |  might be down | Duration: %v\n", link, duration)
		ch <- link
		return
	}
	defer resp.Body.Close() // Always close the response body

	fmt.Printf("%-30s |  %s | Duration: %v\n", link, resp.Status, duration)
	ch <- link
}
