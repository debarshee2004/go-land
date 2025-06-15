package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Pointers package!")

	// var ptr *int
	// fmt.Println("Value of pointer is:", ptr)
	// fmt.Printf("Type of pointer is: %T\n", ptr)

	text := "initial"
	strptr := &text

	// fmt.Println("Value of text is:", text)
	// fmt.Println("Value of pointer to text is:", strptr)
	// fmt.Printf("Type of strptr is: %T\n", strptr)

	doublePnt := &strptr
	fmt.Println("Value of double pointer is:", doublePnt)
	fmt.Println("Value pointed by double pointer is:", *doublePnt)
	fmt.Println("Value pointed by the pointer pointed by double pointer is:", **doublePnt)
	fmt.Println("Value of double pointer is:", doublePnt)
}
