package type_test

import "testing"

type MyInt int64

func TestImpl(t *testing.T){
  var a int32 = 1
  var b int64
  // b = a
  b = int64(a)

  var c MyInt
  // c = b
  c = MyInt(b)

  t.Log(a,b,c)
}

func TestPoint(t *testing.T){
  a:=1
  aPtr:=&a
  aPtr = aPtr+1
  t.Log(a,aPtr)
  t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T){
  var s string
  t.Log("*"+s+"*")
  t.Log(len(s))
  // if s=="" {
  //
  // }
}
