package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to UserInput"
	fmt.Printf(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our work:")

	// comma ok // error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("THanks for rating", input)
}
