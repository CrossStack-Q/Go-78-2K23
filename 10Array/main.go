package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array")

	var fruitList [4]string

	fruitList[0]= "One"
	fruitList[1]= "Two"
	fruitList[3]= "Three"
	

	fmt.Println("Array " , fruitList)
	fmt.Println("Array " , len(fruitList))

	var vegList = [5]string{"Arbi","Tinde","Karela"}
	fmt.Println(vegList)
}