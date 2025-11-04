package main

import "fmt"

// In the main functino, create a slice of ints from 0 through 10.
// Iterate through the slice. For each element, check ot see if the number is even or odd.
//if value is even, print "even". Iff odd print "odd"

func main() {
	start := 1
	end := 11
	step := 1

	numberSlice := make([]int, 0, (end-start)/step)
	for i := start; i < end; i += step {
		numberSlice = append(numberSlice, i)
	} // This starts the slice and appends 1 thorugh 10 to it. To change this values we only need to fix "start" "end" and "step"
	for _, number := range numberSlice {
		fmt.Printf("number %v is %s\n", number, checkEvenOrOdd(number))
	}
}

func checkEvenOrOdd(number int) string {
	if number%2 != 0 {
		return "odd"
	} else {
		return "even"
	}
}
