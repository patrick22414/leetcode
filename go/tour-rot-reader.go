package main

import (
	"fmt"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)

	for i := 0; i < n; i++ {
		ch := p[i]
		if ch >= 'a' && ch <= 'z' {
			p[i] = (ch-'a'+13)%26 + 'a'
		} else if ch >= 'A' && ch <= 'Z' {
			p[i] = (ch-'A'+13)%26 + 'A'
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	fmt.Println(r.r)

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
