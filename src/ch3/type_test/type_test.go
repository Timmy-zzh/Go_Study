package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

func TestImpl(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a
	b = int64(a)

	var c MyInt
	// c = b
	c = MyInt(b)

	fmt.Println(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr+1
	fmt.Println(a, aPtr)
	// fmt.Print("%T %T",a,aPtr)
}

func TestString(t *testing.T) {
	var s string
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
	// if s=="" {
	//
	// }
}
