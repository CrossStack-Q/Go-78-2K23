package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("convertions")
	fmt.Println("Please rate my work ?")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", strings.TrimSpace(input))
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println(numRating + 5)
	if err != nil {
		fmt.Println(err)
	}

}
