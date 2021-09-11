package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	t := make([]byte, len(p))
	rn, rerr := r.r.Read(t)
	if rerr == io.EOF {
		return 0, io.EOF
	}
	for i := 0; i < len(t); i ++ {
		c := t[i]
		if c >= 'a' && c <= 'z' {
			p[i] = (t[i] - 'a' + 13) % 26 + 'a'
		} else if c >= 'A' && c <= 'Z' {
			p[i] = (t[i] - 'A' + 13) % 26 + 'A'
		} else {
			p[i] = t[i]
		}
	}
	return rn, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}