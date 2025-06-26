package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Program on Conversions")
	reader := bufio.NewReader(os.Stdin)

	// Example 1: String to Float Conversion
	fmt.Println("\nExample 1: Convert string input to float")
	fmt.Print("Enter a number rating (1-5): ")
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error converting input to float:", err)
	} else {
		fmt.Println("Your rating is:", numRating)
		fmt.Printf("Type of numRating: %T \n", numRating)
	}

	// Example 2: String to Integer
	fmt.Println("\nExample 2: Convert string to integer")
	str := "42"
	intVal, err := strconv.Atoi(str) // Atoi = ASCII to Integer
	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Printf("String \"%s\" converted to integer: %d\n", str, intVal)
	}

	// Example 3: Integer to String
	fmt.Println("\nExample 3: Convert integer to string")
	age := 21
	ageStr := strconv.Itoa(age) // Itoa = Integer to ASCII
	fmt.Printf("Integer %d converted to string: \"%s\"\n", age, ageStr)

	// Example 4: Float to Integer (Explicit Cast)
	fmt.Println("\nExample 4: Convert float to integer (explicit cast)")
	var floatVal float64 = 3.99
	intCast := int(floatVal) // Truncates decimal
	fmt.Printf("Float %.2f converted to integer: %d\n", floatVal, intCast)

	// Example 5: Integer to Float (Implicit in Go)
	fmt.Println("\nExample 5: Convert int to float (explicit cast)")
	num := 7
	floatConverted := float64(num) // Must cast explicitly
	fmt.Printf("Integer %d converted to float: %.2f\n", num, floatConverted)

	// Example 6: Byte to String and Vice Versa
	fmt.Println("\nExample 6: Byte and string conversions")
	var b byte = 'A'
	s := string(b) // byte to string
	fmt.Printf("Byte %d converted to string: %s\n", b, s)

	str2 := "G"
	byteVal := str2[0] // string to byte
	fmt.Printf("First character of string \"%s\" as byte: %d\n", str2, byteVal)
}
