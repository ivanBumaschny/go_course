package main

// []byte("Hi there!") -> type we want (Value we have)

func main() {
	// decks[startIndexIncliding : uoToNotIncluding]
	// This is how you access a slice of an array

	// hand, remainingCards := Deal(cards, 5)
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
