package main

import "fmt"

// go is a PASS BY VALUE language
// this means that when we pass an object as arguments to a function, a new object in another memory adress is generated and passed as the value of the argument. This means that, if the value we passed is the actual value of the object (in contrast with the pointer of its adress) then a copy of this variable will be generated
// To modify the original variable, we need to pass a pointer as an argument. The new variable will have the value of the adress of the original
// ------ Example ------
// jim.updateName("Jimmy")
// func (p person) updateName() {..}
// Adress	|	Value
// 0000		|
// 0001		|	person{firstName: "Jim"....} <- jim
// 0002		|
// 0003		|	person{firstName: "Jim"....} <- p
//

type person struct {
	firstName   string
	lastName    string
	contactInfo // we can define custom types within a struct
}

type contactInfo struct {
	email   string
	zipCode int
}

// Zero Values:
//
//	Type	|	Zero Value
//	string	|		""
//	int		|		0
//	float	|		0
//	bool	|		false

// ADRESS | VALUE
// Turn ADRESS into VALUE with *adress
// Turn VALUE into ADRESS with &value
func main() {
	// var caitlyn person

	// alex := person{"Alex", "Anderson", contactInfo{"a", 1}}
	// ben := person{firstName: "Ben", lastName: "Brodin"}

	//caitlyn.firstName = "Caitlyn"
	//caitlyn.lastName = "Curry"

	// fmt.Println(alex, ben)
	// fmt.Printf("%+v", caitlyn) // %+v prints all field names and values

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 900,
		},
	}
	// jimPointer := &jim // Create a variable with the address of jim as its value (type pointer)
	// jimPointer.updateName("Jimmy")
	jim.updateName("jimmy") // Go allows to pass the object and interpret it as its pointer. Equivalent to the code above with "jimPointer"
	jim.print()
	fmt.Println("")
	testSliceChange()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (pointerToPerson *person) updateName() { <- The "*pointer" is a type description. It means we are working with a pointer to a person
//
//			*pointerToPerson <- this is an operator. It means we want to manipulate the value the pointer is referencing
//	}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // We are accessing the value of the pointer
}

func testSliceChange() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

// even if you pass by value a slice and change an element inside, it still modifies the original slice, not a supposed copy
//
//			ARRAY				|					SLICES
//	primitive data structure	|			can grow and shrink
//		cant be resized			|	Used 99% of the time for list of elements
//		Rarely used directly	|
//
//	slice contains -> [ptr to head][capacity][length]
func updateSlice(s []string) {
	s[0] = "Bye"
}

// Slice is what Go calls a "REFERENCE TYPE" since its value is just a pointer (reference) to another memory slot
// Passing by value a REFERENCE TYPE still points towards the original memory slot that has the value (not the reference)
// Reference Types: Slices, Maps, Channels, Pointers, Functions <- dont worry about pointers with these
// Value Types: Int, Float, String, Bool, Structs <- use pointers to change these things in a function
