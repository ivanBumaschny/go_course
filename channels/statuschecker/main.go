package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Now we want to, each time a request comes back, to resend it to the same address
	// google.com -> google.com -> google.com -> google.com -> ..
	// amazon.com   ->  amazon.com   ->  amazon.com   ->  amazon.com   -> ..
	// stackoverflow.com -> stackoverflow.com -> stackoverflow.com -> ..
	// for i := 0; i < len(links); i++ {
	// for {
	// go checkLink(<-c, c) // Go expects a string as a first argument, but it is intelligent enough to realize that the channel c is of type string, so it allows it
	// } // Blocking operation
	// Alternative syntax, more readable
	for l := range c {
		// time.Sleep(5 * time.Second) This cannot go here since it is in the main routine, and it pauses it. The main routine should always be prepaired to recieve channels information, without throttle
		// range works as a blocking function. It waits for c to recieve a value, takes it as l and starts the loop. Once finished, it does it again indefinetly.
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	// We could paste the  time.Sleep(5 * time.Second) here, but it would not be propper syntax. That is why we use a lambda/function literal
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
