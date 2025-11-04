package main

import "fmt"

// cannot instance values of an interface type
type bot interface {
	// If you are a type in this program, that has a getGreeting reciever function that returns a string, you are an honorary member of the type "bot"
	getGreeting() string
}

// Interfeces are IMPLICIT in Go because we do not need to define any link between the types and the interface that interacts with them (ex. bot-englishBot). Go takes care of it for us
// type englishBot implements bot struct{} <- This would b explicit interface, which is not correct in go

// Go does not support overloading (multiple functions with same name in same runtime)

type englishBot struct{}
type spanishBot struct{}

// Interfaces lets us reuse functions that can be applied to many different types of structures without the need of writting the reciever function for each structure type
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// If we are not using the self varaible (eb in this case) we can delete it
func (englishBot) getGreeting() string {
	// very different logic for generating an english greeting (than spanish)
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// very different logic for generating a spanish greeting (than english)
	return "Hola!"
}
