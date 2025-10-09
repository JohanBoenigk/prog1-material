package guessinggame

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	var n int
	fmt.Print("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

func guessinggame() {
	my_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if NumberGood(guess, my_number) {
			fmt.Println("Richtig geraten! :-)")
			return
		} else {
			fmt.Println("Falsch geraten :)")
		}
	}
	fmt.Println("Zu viele flasche Zahlen! :-(")
}

func NumberGood(g, e int) bool {
	return g == e
}

func main() {
	guessinggame()
}
