package main

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(x []byte) (int, error) {
	n, err := rot.r.Read(x)
	for i := 0; i < n; i++ {
		if ('a' <= x[i] && x[i] + 13 <= 'z' || 'A' <= x[i] && x[i] + 13 <= 'Z') {
			x[i] += 13
		} else if ('a' <= x[i] - 13 && x[i] <= 'z' || 'A' <= x[i] -13 && x[i] <= 'Z') {
			x[i] -= 13
		}
	}
	fmt.Println(n)
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
