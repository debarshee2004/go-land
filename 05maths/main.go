package main

import (
	"fmt"
	"math"
	"math/big"

	"crypto/rand"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to the Maths package!")

	// -----------------------
	// Example 1: Basic math operations
	fmt.Println("\nExample 1: Math operations")
	x := 25.0
	y := 4.0

	fmt.Printf("Sqrt(%.1f) = %.2f\n", x, math.Sqrt(x))
	fmt.Printf("Pow(%.1f, %.1f) = %.2f\n", x, y, math.Pow(x, y))
	fmt.Printf("Ceil(%.2f) = %.2f\n", 2.3, math.Ceil(2.3))
	fmt.Printf("Floor(%.2f) = %.2f\n", 2.8, math.Floor(2.8))
	fmt.Printf("Abs(-%.2f) = %.2f\n", x, math.Abs(-x))
	fmt.Printf("Max(%.2f, %.2f) = %.2f\n", x, y, math.Max(x, y))
	fmt.Printf("Min(%.2f, %.2f) = %.2f\n", x, y, math.Min(x, y))

	// -----------------------
	// Example 2: Trigonometric functions
	fmt.Println("\nExample 2: Trigonometry")
	angle := 30.0
	radian := angle * (math.Pi / 180) // Degrees to radians

	fmt.Printf("Sin(%.1f°) = %.4f\n", angle, math.Sin(radian))
	fmt.Printf("Cos(%.1f°) = %.4f\n", angle, math.Cos(radian))
	fmt.Printf("Tan(%.1f°) = %.4f\n", angle, math.Tan(radian))

	// -----------------------
	// Example 3: Generate random number using math/rand
	// fmt.Println("\nExample 3: Pseudo-random number with math/rand")
	// mathrand.Seed(time.Now().UnixNano()) // Seed with current time
	// randomInt := mathrand.Intn(10)       // Generates number in [0, 9]
	// fmt.Printf("math/rand Intn(10): %d\n", randomInt)

	// randomFloat := mathrand.Float64() // Float between 0.0 and 1.0
	// fmt.Printf("math/rand Float64(): %.4f\n", randomFloat)

	// -----------------------
	// Example 4: Secure random number using crypto/rand
	fmt.Println("\nExample 4: Secure random number with crypto/rand")
	secureRand, err := rand.Int(rand.Reader, big.NewInt(10)) // [0,9]
	if err != nil {
		fmt.Println("Error generating secure random number:", err)
	} else {
		fmt.Println("crypto/rand Int (0-9):", secureRand)
		fmt.Printf("Type of secureRand: %T\n", secureRand)
	}
}
