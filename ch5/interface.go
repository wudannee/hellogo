package ch5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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

func DemoInterface() {
	x := &X{true}
	io.Copy(os.Stdout, x)

	foo1()
	foo2()
}

func foo1() {
	newWriter1 := bufio.NewWriter(os.Stdout)
	newWriter1.WriteString("bufio writer")
	newWriter1.Flush()

	fmt.Println("\n---")
	strBuilder1 := &strings.Builder{}
	strBuilder1.WriteString("string builder 1")
	strBuilder1.WriteString("\n string builder 2")
	fmt.Fprintln(strBuilder1, "\n string builder 3")
	fmt.Println("buf:", strBuilder1.String())
}

type Notifier interface {
	Notify()
}

type Alarm struct{}

func (a *Alarm) Notify() {
	fmt.Println("alarm is ringing")
}

type Phone struct{}

func (p *Phone) Notify() {
	fmt.Println("phone is ringing")
}

func foo2() {
	var notifiers = []Notifier{
		&Alarm{},
		&Phone{},

		// will not compile because Notify has a pointer as its receiver
		// Phone{},
	}

	notify := func(n Notifier) {
		n.Notify()
	}

	for _, n := range notifiers {
		notify(n)
	}
}
