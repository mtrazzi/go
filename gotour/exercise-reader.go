package main

import (
	"golang.org/x/tour/reader"
	"strings"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(x []byte) (int, error) {
	r := strings.NewReader("A")
	n, err := r.Read(x)
	return n, err
}

func main() {
	reader.Validate(MyReader{})
}
