package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to input program code"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Comma ok syntax OR Comma err syntax
	fmt.Println("You entered:", input)
}
