func guessinggame() {
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if Numbergood(guess) {
			fmt.Println("Richtig geraten! :-)")
			return
		}
	}
	fmt.Println("Zu viele flasche Zahlen! :-(")
}