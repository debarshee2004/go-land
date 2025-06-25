package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Go Functions Example!")
	result := add(5, 3)
	fmt.Println("The result is:", result)

	fmt.Println("The result of proAdder is:", proAdder(1, 2, 3, 4, 5))
	fmt.Println("The result of proAdder with no arguments is:", proAdder())

	a, b := doubleReturn()
	fmt.Println("The two returned values are:", a, "and", b)
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func doubleReturn() (int, int) {
	return 5, 10
}
