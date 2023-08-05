package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Mango", "Pineapple "}
	fmt.Printf("Type of FruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Orange")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 13
	highScores[2] = 3
	highScores[3] = 12

	highScores = append(highScores, 122, 32, 323)

	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

}
