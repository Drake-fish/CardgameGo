package main

func main() {
	cards := newDeck()

	// 	hand, remainingDeck := deal(cards, 5)

	// 	hand.print()
	// 	remainingDeck.print()
	// }

	// func newCard() string {
	// 	return "Five of Diamonds"
	// cards.saveToFile("my_cards")
	// cards2 := newDeckFromFile("my_cards")

	// cards2.print()

	cards.shuffle()
	cards.print()
}
