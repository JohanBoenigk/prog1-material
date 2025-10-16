package main

import (
	"fmt"
	"math/rand"
)

func main() {
	my_number := rand.Intn(100) + 1
	for i := 0; i < 3; i = i + 1 {
		guess := RReadNumber()
		if NumberGood(my_number, guess) {
			fmt.Println("Das war Korrekt!")
			return
		}
		fmt.Printf("%d ist nicht korrekt!\n", guess)
	}
	fmt.Println("Game Over")
}
func RReadNumber() int {
	fmt.Print("Rate ein Zahl: ")
	var n int
	fmt.Scan(&n)
	return n
}
func NumberGood(x int, n int) bool {
	return x == n
}
