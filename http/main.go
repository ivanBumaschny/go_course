package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com") // Generates a *http.Response value type
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(resp.Body)
	bs := make([]byte, 99999) // Give me an empty slice with 99999 elements
	resp.Body.Read(bs)        // this saves the response body into the byteslice

	// the resp.Body attribute of the resp struct is a ReadClose interface. This interface implements two inner interfaces that need to be (at least one of them) satisfied in order to consider the value type a ReadClose. The Read inner interface has a function Read([]byte) defined inside that reads a byteslice. This means that the resp.Body attribute itself can call the Read([]byte) function
	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// This is the best way to do this
	// io.Copy(os.Stdout, resp.Body)

	// type Writer interface {Writer(p []byte) (n int, err error)}
	// th function Copy expects a type objects that implements the Writer interface. The second argument needs to implement the Reader interface
	// resp.Body -> Implements the Reader interface
	// os.Stdout -> value of type File -> File has a function called Write (therefore, it implements the Writer interface and can be considered a Writer)
}

// we can define a complex interface which has interfaces inside.
// this means that you need to satisfy the inner interfaces to define a type that can be considered the complex type that you defined

func (logWriter) Write(bs []byte) (int, error) {
	// The logWriter now implements the Writer interface
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
