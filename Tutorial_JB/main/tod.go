package main

import (
	"fmt"
)

func main() {
	var age int
	var name string
	fmt.Print("Bitte gib dein Alter ein: ")
	fmt.Scan(&age)
	fmt.Print("Bitte gibt deinen Namen ein: ")
	fmt.Scan(&name)
	fmt.Printf("Hallo deine Name ist %v und du bist %v Jahre alt.\n", name, age)

	if age < 18 {
		var missing = 18 - age
		if missing == 1 {
			fmt.Printf("Du bist nicht Volljährig. Du brauchst noch %v Jahr, bis du volljährig bist.", missing)
		} else {
			fmt.Printf("Du bist nicht Volljährig. Du brauchst noch %v Jahre, bis du volljährig bist.", missing)
		}
	} else {
		fmt.Println("Du bist Volljährig")
	}
}
