package constant_test

import (
  "fmt"
  "testing"
)

const(
  Monday = iota+1
  Tuesday
  Wenday
)

const(
  Readable = 1<<iota
  Writable
  Executable
)

func TestConstTry(t *testing.T){
  fmt.Print(Monday,Tuesday)
  fmt.Println()
}

func TestConstTry1(t *testing.T){
  a:=7 //0111
  fmt.Print(a&Readable==Readable,a&Writable==Writable)
  fmt.Println()
}
