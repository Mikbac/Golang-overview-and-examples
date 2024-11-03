package main

func runFunctionsDemo() {
	card := getCard()
	card1, card2, cardQuantity := getCards()

	println(card)
	println(card1, card2, cardQuantity)
}

func getCard() string {
	return "sample card"
}

func getCards() (string, string, int) {
	return "sample card 1", "sample card 2", 2
}
