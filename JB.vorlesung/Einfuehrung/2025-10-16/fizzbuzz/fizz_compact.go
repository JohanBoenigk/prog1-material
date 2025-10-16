package fizzbuzz

import (
	"fmt"
)

func Fizzcompact() {
	for i := 1; i <= 30; i++ { // hier wird bis 30 gezählt
		// wenn i weder durch 3 noch durch 5 teilbar ist
		if i%3 != 0 && i%5 != 0 { // wenn ich durch 3 oder 5 teilen kann hab ich einen Rest von 0
			fmt.Println(i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
		}
		fmt.Println()
	}
}
