package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(stream []byte) (int, error) {
	for x := 0; x < 12; x++ {
		//fmt.Println(cap(stream), x)
		stream[x] = 'A'
	}
	return 12, nil
}

func main() {
	reader.Validate(MyReader{})
}
