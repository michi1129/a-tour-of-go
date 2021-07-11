package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(buf []byte) (int, error) {
	l, e := r.r.Read(buf)
	if e != nil {
		return l, e
	}
	for i := 0; i < l; i++ {
		if buf[i] >= 'A' && buf[i] <= 'M' || buf[i] >= 'a' && buf[i] <= 'm' {
			buf[i] += 13
		} else if buf[i] >= 'N' && buf[i] <= 'Z' || buf[i] >= 'n' && buf[i] <= 'z' {
			buf[i] -= 13
		}
	}
	return l, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
