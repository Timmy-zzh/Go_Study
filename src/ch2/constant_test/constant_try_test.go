package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wenday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstTry(t *testing.T) {
	fmt.Println(Monday, Tuesday)
}

func TestConstTry1(t *testing.T) {
	a := 7 //0111
	fmt.Println(a&Readable == Readable, a&Writable == Writable)
}
