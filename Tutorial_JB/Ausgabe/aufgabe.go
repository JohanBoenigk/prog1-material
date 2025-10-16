package ausgabe

import (
	"fmt"
	"strings"
)

func Foo1() {
	x := 8
	y := 7
	k := x + 2*y - 2

	fmt.Println(x, y, k)
	fmt.Printf("%d-%d-%d\n", x, y, k)
	fmt.Println(k + 2 - x)
	fmt.Println(float64(k) / float64(x))
}
func Foo2(n int) {
	s := "*"
	for i := 0; i < n; i++ {
		fmt.Println(s)
		s += s
	}
}
func Foo3(n int) {
	if n == 0 {
		return
	}
	Foo3(n - 1)
	fmt.Println(strings.Repeat("*", n))
}

//func main() {
//z := 5
//Foo1()
//Foo2(z)
//Foo3(z)
//}
