package main

func main() {
	// cards := newDeck()

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	cardsFromFile := newDeckFromFile("my_cards")

	cardsFromFile.print()
}
