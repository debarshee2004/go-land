package main

import (
	"fmt"
	"sync"
	"time"
)

// Global variables for demonstration
var counter int
var balance int = 1000

func main() {
	// Example 1: Race condition demonstration (without mutex)
	fmt.Println("1. Race condition without mutex:")
	demonstrateRaceCondition()

	// Example 2: Fixed with Mutex
	fmt.Println("\n2. Race condition fixed with mutex:")
	demonstrateWithMutex()

	// Example 3: RWMutex example
	fmt.Println("\n3. RWMutex example (multiple readers, single writer):")
	demonstrateRWMutex()

	// Example 4: Bank account simulation with mutex
	fmt.Println("\n4. Bank account simulation with mutex:")
	bankAccountExample()

	// Example 5: Original example - fixed
	fmt.Println("\n5. Original example - corrected:")
	originalExampleFixed()
}

// Example 1: Race condition without mutex
func demonstrateRaceCondition() {
	counter = 0
	wg := &sync.WaitGroup{}

	// Start 1000 goroutines that increment counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			counter++ // Race condition here!
		}(wg)
	}

	wg.Wait()
	fmt.Printf("   Counter value (should be 1000): %d\n", counter)
}

// Example 2: Fixed with mutex
func demonstrateWithMutex() {
	counter = 0
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// Start 1000 goroutines that increment counter safely
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex) {
			defer wg.Done()
			m.Lock()
			counter++ // Safe increment
			m.Unlock()
		}(wg, mut)
	}

	wg.Wait()
	fmt.Printf("   Counter value (should be 1000): %d\n", counter)
}

// Example 3: RWMutex demonstration
func demonstrateRWMutex() {
	data := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	rwMut := &sync.RWMutex{}

	// Start multiple readers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			m.RLock() // Read lock - multiple readers allowed
			fmt.Printf("   Reader %d: %v\n", id, data)
			time.Sleep(100 * time.Millisecond)
			m.RUnlock()
		}(i, wg, rwMut)
	}

	// Start a writer
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond) // Let readers start first
		m.Lock()                          // Write lock - exclusive access
		data = append(data, 6)
		fmt.Printf("   Writer: Added element, new data: %v\n", data)
		m.Unlock()
	}(wg, rwMut)

	wg.Wait()
}

// Example 4: Bank account simulation
func bankAccountExample() {
	balance = 1000
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// Simulate concurrent transactions
	transactions := []struct {
		name   string
		amount int
	}{
		{"Deposit", 200},
		{"Withdrawal", 150},
		{"Deposit", 300},
		{"Withdrawal", 100},
		{"Withdrawal", 50},
	}

	for _, tx := range transactions {
		wg.Add(1)
		go func(name string, amount int, wg *sync.WaitGroup, m *sync.Mutex) {
			defer wg.Done()

			m.Lock()
			oldBalance := balance
			if name == "Deposit" {
				balance += amount
				fmt.Printf("   %s: +$%d, Balance: $%d -> $%d\n", name, amount, oldBalance, balance)
			} else {
				if balance >= amount {
					balance -= amount
					fmt.Printf("   %s: -$%d, Balance: $%d -> $%d\n", name, amount, oldBalance, balance)
				} else {
					fmt.Printf("   %s: -$%d FAILED (insufficient funds), Balance: $%d\n", name, amount, balance)
				}
			}
			m.Unlock()

			time.Sleep(10 * time.Millisecond) // Simulate processing time
		}(tx.name, tx.amount, wg, mut)
	}

	wg.Wait()
	fmt.Printf("   Final balance: $%d\n", balance)
}

// Example 5: Original example - corrected
func originalExampleFixed() {
	fmt.Println("   Race condition fixed - LearnCodeonline.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4) // Fixed: Add correct number of goroutines

	// Writer goroutines
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("   Writer One")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("   Writer Two")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("   Writer Three")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
	}(wg, mut)

	// Reader goroutine
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond) // Let writers execute first
		fmt.Println("   Reader")
		m.RLock()
		fmt.Printf("   Current score: %v\n", score)
		m.RUnlock()
	}(wg, mut)

	wg.Wait()
	fmt.Printf("   Final score: %v\n", score)
}
