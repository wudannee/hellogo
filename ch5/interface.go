package ch5

import (
	"io"
	"os"
)

func DemoInterface() {
	x := &X{true}
	io.Copy(os.Stdout, x)

}

type X struct {
	NewLine bool
}

func (x *X) Read(p []byte) (n int, err error) {
	s := "hello x"
	if x.NewLine {
		s += "\n"
	}
	n = copy(p, s)
	return n, io.EOF
}

func (x *X) Write(p []byte) (n int, err error) {
	return 0, io.EOF
}
