package main

import "fmt"

func main() {
	fmt.Println("Now() allocates memory but not initialize it . Memory address is there . Zeroed Storage .")
	fmt.Println("Make() allocates memory and initialize it . Memory address is there . Non-Zeroed Storage .")
	fmt.Println("GC(Garbage collection) happens automatically in it . Out of scope is nil . GOGC set initial garbage collection percentage . ")
	
}