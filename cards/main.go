package main

func main() {
	// cards := newDeck()
	// cards = append(cards, "Six of Spades")
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi There"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
