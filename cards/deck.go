package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever to indicate that the function corresponds to Deck
// The actual copy of the Deck we are working with is available in the function as a variable called 'd'
// very variable of type Deck can call this function
func (d Deck) print() {
	// d is equivalent of using "self" in other languages
	for i, card := range d {
		fmt.Println(i, card)
		// index, card -> index of element, element
		// range cards -> slice of "cards" to iterate over
		// ":=" in for loops, we need to redeclare each loop
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// helper function
func (d Deck) toString() string {
	return strings.Join([]string(d), ",") // concatenates all items in a string array into a single string, each item separated by a comma
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) // anyone can read and write
}

func newDeckFromFile(filename string) Deck {
	// bs -> byteslice
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)

}

func (d Deck) shuffle() { // we dont return a deck, just shuffle the deck in question
	for i := range d {
		newPosition := rand.IntN(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // we dont need a third variable in Go, just switch it like this
	}
}
