package constant_test

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
  t.Log(Monday,Tuesday)
}

func TestConstTry1(t *testing.T){
  a:=7 //0111
  t.Log(a&Readable==Readable,a&Writable==Writable)
}
