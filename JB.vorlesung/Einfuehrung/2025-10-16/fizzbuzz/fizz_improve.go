package fizzbuzz

import (
	"fmt"
)

func FizzImproved(x, y, n int) {
	for i := 1; i <= n; i++ { // hier wird von 1 bis n gezählt
		// wenn i durch x und y teilbar ist, gib "fizzbuzz" aus
		if i%x == 0 && i%y == 0 {
			fmt.Println("FizzBuzz")
		} else

		// wenn i durch x teilbar ist, gib "fizz" aus
		if i%x == 0 {
			//fmt.Println(i, "fizz")
			//fmt.Println("%10v: %v\n", i, "Fizz")
			fmt.Println("Fizz")
		} else
		// wenn i durch y tielbar ist gib "buzz" aus
		if i%y == 0 {
			fmt.Println("Buzz")
			// sonst gibt i aus
		} else {
			fmt.Println(i)
		}

	}
}
