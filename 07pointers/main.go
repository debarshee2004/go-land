package main

import (
	"fmt"
)

// Topic 06 - Pointers in Go

func main() {
	fmt.Println("Welcome to the Pointers package!")

	// ----------------------------
	// Example 1: Declaring a nil pointer
	fmt.Println("\nExample 1: Nil pointer declaration")
	var ptr *int                                        // Declares a pointer to an int, but it points to nothing (nil)
	fmt.Println("Value of uninitialized pointer:", ptr) // nil
	fmt.Printf("Type of ptr: %T\n", ptr)

	// ----------------------------
	// Example 2: Using pointer to get and set values
	fmt.Println("\nExample 2: Assigning address to pointer")
	value := 100
	ptr = &value // Assigns the address of value to ptr
	fmt.Println("Address stored in ptr:", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr) // Dereferencing

	*ptr = 200 // Changing value at the address
	fmt.Println("Updated value through ptr:", value)

	// ----------------------------
	// Example 3: Pointers with strings
	fmt.Println("\nExample 3: Pointer to a string")
	text := "initial"
	strptr := &text
	fmt.Println("Value of text:", text)
	fmt.Println("Address of text (strptr):", strptr)
	fmt.Println("Value at address of strptr:", *strptr)

	// ----------------------------
	// Example 4: Double pointer (pointer to a pointer)
	fmt.Println("\nExample 4: Double pointer")
	doublePnt := &strptr // strptr is *string, so &strptr is **string
	fmt.Println("doublePnt (type **string):", doublePnt)
	fmt.Println("Value pointed by doublePnt (*doublePnt):", *doublePnt)       // value of strptr
	fmt.Println("Value pointed by *(*doublePnt) (**doublePnt):", **doublePnt) // value of text

	// ----------------------------
	// Example 5: Function with pointer parameter
	fmt.Println("\nExample 5: Passing pointer to function")
	num := 5
	fmt.Println("Before function call, num =", num)
	increment(&num) // pass pointer
	fmt.Println("After function call, num =", num)

	// ----------------------------
	// Example 6: Returning pointer from function
	fmt.Println("\nExample 6: Returning pointer from function")
	ptrFromFunc := newInt()
	fmt.Println("Value from returned pointer:", *ptrFromFunc)
}

// ----------------------------
// Function that modifies value through pointer
func increment(n *int) {
	*n = *n + 1
}

// Function that returns a pointer to an int
func newInt() *int {
	i := 42
	return &i
}
