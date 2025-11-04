package main // MAPS are like DICTS in python
// All keys within a map should b the sme type, and all values of those keys too
import "fmt"

func main() {
	// var colors map[string]string <- these two are equivalent
	// colors := make(map[string]string) <- these two are equivalent

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["yellow"] = "#ffffff" // The key should be the type defined (in this case string)
	// structName.white <- this does not work with maps

	delete(colors, "yellow") // This deletes an entry of a map by the key

	colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		// color -> key for this step through the loop
		// hex -> value for this step through the loop
		fmt.Println("Hex code for", color, "is", hex)
	}
}

//		MAP
// All keys must be the same type
// All values must be the same type
// Use to represent a collection of relatd properties
// Dont need to know all the keys at compile time
// Keys are indexed - we can iterate over them
// REFERENCE TYPE

//		STRUCT
// All keys can be of different type
// All values can be of different type
// You need to know all the different fields at compile time
// Use to represent a "thing" with a lot of different properties
// VALUE TYPE
