package main

import (
	"fmt"
	"math/big"

	"crypto/rand"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to the Maths package!")

	// random number using math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number using crypto/rand
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Random number between 0 and 9:", randomNumber)
	fmt.Printf("Random number type: %T\n", randomNumber)
}
