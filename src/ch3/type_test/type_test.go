package type_test

import (
  "fmt"
  "testing"
)

type MyInt int64

func TestImpl(t *testing.T){
  var a int32 = 1
  var b int64
  // b = a
  b = int64(a)

  var c MyInt
  // c = b
  c = MyInt(b)

  fmt.Print(a,b,c)
  fmt.Println()
}

func TestPoint(t *testing.T){
  a:=1
  aPtr:=&a
  // aPtr = aPtr+1
  fmt.Print(a,aPtr)
  // fmt.Print("%T %T",a,aPtr)
    fmt.Println()
}

func TestString(t *testing.T){
  var s string
  fmt.Print("*"+s+"*")
  fmt.Println()
  fmt.Print(len(s))
  // if s=="" {
  //
  // }
  fmt.Println()
}
