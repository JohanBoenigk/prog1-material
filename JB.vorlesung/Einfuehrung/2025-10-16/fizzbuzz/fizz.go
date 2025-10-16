package fizzbuzz

import (
	"fmt"
)

func Fizz() {
	for i := 1; i <= 30; i++ { // hier wird bis 30 gezählt
		// wenn i durch 3 und 5 teilbar ist, gib "fizzbuzz" aus
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else

		// wenn i durch 3 teilbar ist, gib "fizz" aus
		if i%3 == 0 {
			//fmt.Println(i, "fizz")
			//fmt.Println("%10v: %v\n", i, "Fizz")
			fmt.Println("Fizz")
		} else
		// wenn i durch 5 tielbar ist gib "buzz" aus
		if i%5 == 0 {
			fmt.Println("Buzz")
			// sonst gibt i aus
		} else {
			fmt.Println(i)
		}

	}
}
