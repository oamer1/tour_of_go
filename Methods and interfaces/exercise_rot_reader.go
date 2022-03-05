package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)

	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err

}

// https://en.wikipedia.org/wiki/ROT13#Go
func rot13(b byte) byte {
	//ASCII 65 is A and 90 is Z
	var shift byte = 13
	if b >= 'A' && b <= 'Z' {

		base := b - 'A'
		b = ((base + shift) % 26) + 65

	} else if b >= 'a' && b < 'z' {
		//ASCII 97 is a and 122 is z

		base := b - 'a'
		b = ((base + shift) % 26) + 97

	}

	return b

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
