package main

import (
	"fmt"
	"io"
	"os"
)

//encode/decode - directly change to and from JSON to Go Data Structure

//writer interface
//method Write(p []byte) (n int, err error)
//any other type that implements the Write funciton will also be of type Writer

func main() {
	fmt.Fprintln(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "Hello")
}
