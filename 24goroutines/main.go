package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// Example 1: Basic goroutine with anonymous function
	fmt.Println("1. Anonymous function goroutine:")
	go func() {
		fmt.Println("   Anonymous goroutine executed!")
	}()
	time.Sleep(100 * time.Millisecond) // Wait for goroutine to complete

	// Example 2: Goroutine with parameters
	fmt.Println("\n2. Parameterized goroutines:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("   Goroutine %d is running\n", id)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// Example 3: Using WaitGroup with simple tasks
	fmt.Println("\n3. WaitGroup example with simple tasks:")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()

	// Example 4: Goroutines with channels
	fmt.Println("\n4. Goroutines with channels:")
	ch := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			ch <- fmt.Sprintf("Message from goroutine %d", id)
		}(i)
	}

	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Printf("   Received: %s\n", msg)
	}

	// Example 5: Original website status checker
	fmt.Println("\n5. Website status checker with goroutines:")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Printf("   Processed websites: %v\n", signals)

	// Example 6: Goroutines with timeout
	fmt.Println("\n6. Goroutine with timeout pattern:")
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("   Long-running task completed!")
	case <-time.After(1 * time.Second):
		fmt.Println("   Task timed out!")
	}

	fmt.Println("\n=== All goroutine examples completed ===")
}

// Worker function for WaitGroup example
func worker(id int) {
	defer wg.Done()
	fmt.Printf("   Worker %d starting\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Printf("   Worker %d finished\n", id)
}

// Greeter function - restored and improved
func greeter(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("   %s - iteration %d\n", s, i+1)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("   Error accessing %s: %v\n", endpoint, err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		res.Body.Close() // Good practice to close response body
		fmt.Printf("   %d status code for %s\n", res.StatusCode, endpoint)
	}
}
