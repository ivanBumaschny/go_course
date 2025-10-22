package main

// func main() {
// 	// the third word is telling the compiler that only a string will be assigned to the compiler
// 	// var card string = "Ace of Spades"
// 	//card := "Ace of Spades" // Alternative version
// 	cards := deck{"Ace of Diamonds", newCard()} // declare a slice of type string
// 	cards = append(cards, "Six of Spades")

// 	// this is the equivalent of doing: func print(deck) -> this means that we pass this cards object to the print function, which is define to only recieve this type of structures
// 	cards.print()
// }

func main() {
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
