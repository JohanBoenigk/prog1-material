package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//my_number := 42
	// NumberGood prüft, ob x gleich einer zufällig gewählten Zahl zwischen 1 und 100 ist.
	// Liefert true, falls x gleich dieser zufallszahl ist ansonsten false.
	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i = i + 1 {
		guess := ReadNumber()
		// if guess == my_number {
		if guess == my_number {
			fmt.Println("Das war richtig!")
			return
		}
		if guess < my_number {
			fmt.Println("Die gesuchte Zahl ist größer.")
		} else {
			// if guess > my_number { ist auch möglich
			fmt.Println("Die gesuchte Zahl ist kleiner.")
		}
	}
	fmt.Println("Game Over!")
}

// ReadNumber liefert und ein int (Zahl).
func ReadNumber() int {
	var n int

	fmt.Print("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)

	return n
}
