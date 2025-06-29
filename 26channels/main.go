package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Example 1: Basic unbuffered channel
	fmt.Println("1. Basic unbuffered channel:")
	basicChannel()

	// Example 2: Buffered channel
	fmt.Println("\n2. Buffered channel:")
	bufferedChannel()

	// Example 3: Channel directions (send-only, receive-only)
	fmt.Println("\n3. Channel directions (send-only, receive-only):")
	channelDirections()

	// Example 4: Channel closing and range
	fmt.Println("\n4. Channel closing and range:")
	channelClosing()

	// Example 5: Select statement with channels
	fmt.Println("\n5. Select statement with channels:")
	selectExample()

	// Example 6: Worker pool pattern
	fmt.Println("\n6. Worker pool pattern:")
	workerPool()

	// Example 7: Original example - improved
	fmt.Println("\n7. Original example - improved:")
	originalExampleImproved()
}

// Example 1: Basic unbuffered channel
func basicChannel() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	message := <-ch
	fmt.Printf("   Received: %s\n", message)
}

// Example 2: Buffered channel
func bufferedChannel() {
	ch := make(chan int, 3) // Buffer size of 3

	// Send values without blocking
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("   Channel length: %d, capacity: %d\n", len(ch), cap(ch))

	// Receive values
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("   Received: %d\n", val)
	}
}

// Example 3: Channel directions
func channelDirections() {
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Send-only channel
	go sender(ch, wg)

	// Receive-only channel
	go receiver(ch, wg)

	wg.Wait()
}

func sender(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("   Sender: Sending values...")

	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("   Sender: Sent %d\n", i)
	}
	close(ch)
}

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond) // Let sender start first

	fmt.Println("   Receiver: Receiving values...")
	for val := range ch {
		fmt.Printf("   Receiver: Got %d\n", val)
	}
}

// Example 4: Channel closing and range
func channelClosing() {
	ch := make(chan int, 5)

	// Send values
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * i
		}
		close(ch) // Important: close the channel
	}()

	// Receive using range (automatically stops when channel is closed)
	fmt.Println("   Squares:")
	for val := range ch {
		fmt.Printf("   %d ", val)
	}
	fmt.Println()

	// Check if channel is closed
	val, ok := <-ch
	fmt.Printf("   Channel closed check - Value: %d, Open: %t\n", val, ok)
}

// Example 5: Select statement
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select waits for multiple channel operations
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   Received %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   Received %s\n", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("   Timeout!")
		}
	}
}

// Example 6: Worker pool pattern
func workerPool() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	fmt.Println("   Results:")
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("   Job result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond) // Simulate work
		results <- job * 2                // Send result
	}
}

// Example 7: Original example - improved
func originalExampleImproved() {
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Receive-only goroutine
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		val, isChannelOpen := <-myCh
		fmt.Printf("   Receiver: Value=%d, Channel Open=%t\n", val, isChannelOpen)

		// Try to receive again after channel is closed
		if val2, ok := <-myCh; ok {
			fmt.Printf("   Receiver: Second value=%d\n", val2)
		} else {
			fmt.Printf("   Receiver: Channel is closed, got zero value: %d\n", val2)
		}
	}(myCh, wg)

	// Send-only goroutine
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("   Sender: Sending value 42")
		myCh <- 42

		fmt.Println("   Sender: Sending value 100")
		myCh <- 100

		fmt.Println("   Sender: Closing channel")
		close(myCh)
	}(myCh, wg)

	wg.Wait()
	fmt.Println("   Original example completed")
}
