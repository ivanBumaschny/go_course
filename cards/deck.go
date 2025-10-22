package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever to indicate that the function corresponds to deck
// The actual copy of the deck we are working with is available in the function as a variable called 'd'
// very variable of type deck can call this function
func (d deck) print() {
	// d is equivalent of using "self" in other languages
	for i, card := range d {
		fmt.Println(i, card)
		// index, card -> index of element, element
		// range cards -> slice of "cards" to iterate over
		// ":=" in for loops, we need to redeclare each loop
	}
}
