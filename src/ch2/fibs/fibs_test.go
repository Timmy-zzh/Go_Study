package fibs

import (
	"fmt"
	"testing"
)

func TestFibsList(t *testing.T) {
	var a int = 1
	var b int = 2
	var c int = 765
	// t.Log(a)
	// t.Log(a,b,c)
	fmt.Println(a, b, c)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
	fmt.Println(a, b)
}
