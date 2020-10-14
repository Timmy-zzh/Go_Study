package loop

import (
	"fmt"
	"testing"
)

//for 循环语句学习
func TestLoop(t *testing.T){
  var a int = 1
  b:=2
  fmt.Println(a,b)

  n:=0
  for n<5 {
    fmt.Println(n)
    n++
  }

}
