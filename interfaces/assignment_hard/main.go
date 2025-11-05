package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filename := os.Args[1]
	fmt.Println("File name opened:", filename)

	// *os.File has this reciever func
	// func (f *os.File) Read(b []byte) (n int, err error)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file", filename)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	lw := logWriter{}
	io.Copy(lw, file) // the *File typ implements a Reader interface
}

func (logWriter) Write(bs []byte) (int, error) {
	// The logWriter now implements the Writer interface
	fmt.Println("Content of the file:", string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
