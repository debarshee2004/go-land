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

	fmt.Println("Enter a number rating (1-5):")
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to float:", err)
	} else {
		fmt.Println("Your rating is:", numRating)
		fmt.Printf("Type of numRating: %T \n", numRating)
	}
}
