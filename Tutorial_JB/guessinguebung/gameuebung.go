package main

import (
	"fmt"
	"math/rand"
)

func main() {
	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Das ist die Richtige Zahl")
			return
		}
		fmt.Println("Das ist die falsche Zahl")
	}
	fmt.Println("Game Over!")
}

func ReadNumber() int {
	var n int

	fmt.Println("Gib eine beliebige Zahl ein: ")
	fmt.Scan(&n)

	return n
}

func NumberGood(x int, n int) bool {
	return x == n
}
