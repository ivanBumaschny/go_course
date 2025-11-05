package main

import (
	"fmt"
	"net/http"
)

// main() function instantiates the "Main Routine" (only one)
func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	// Channel connects all routines to one another (regardless of which a routine is a Main or a Child routine)
	// Channels are Typed, meant that each channel only shares data of that type. One Channel per type that needs to be shared
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // The "go" command instantiates "Child Routines" (can be many of them)
	}
	fmt.Println(<-c) // Wait for a value to be sent through the channel. When we get one, log it out immediately
	// this line is a blocking line. It waits for any routine to finish and pass a message.

	for i := 0; i < len(links); i++ {
		// The code loads the channel, then enters this for loop. Once inside, the messages start arriving. Since the Println() is a block call, each loop prints one of the messages, each time they arrive. If they arrive before the next iteration, a stack is created and handled
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up!"
}
