package main

import (
	"fmt"
	"prog1-material/jb.vorlesung/greet"
)

func main() {
	fmt.Println(greet.Greet("Kurs"))
	fmt.Println("Alle " + greet.Greet("Kursteilnehmer"))
}
