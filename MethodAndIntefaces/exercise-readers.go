package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(p []byte) (n int, err error)  {
	c := len(p)
	for i := 0; i < c; i ++ {
		p[i] = 'A'
	}
	return c, nil
}

func main() {
	reader.Validate(MyReader{})
}