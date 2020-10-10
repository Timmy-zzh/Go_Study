package fib

import (
  "fmt"
  "testing"
)

func TestFibList(t *testing.T){
  //赋值方式一
  var a int = 1
  var b int = 1

// //赋值方式二
//   var (
//     a =1
//     b =2
//   )
//
// //赋值方式三
//   a:=1
//   b:=2

  // fmt.Print(a,b)

  for i:= 0; i < 5; i++ {
    tmp:=a
    a = b
    b = tmp+a
     fmt.Print(b ," ")
  }
  fmt.Println()
}

//两个变量交换
func TestExchange(t *testing.T){
  a:=1
  b:=2

  // temp:=a
  // a =  b
  // b = temp

  a,b = b,a
  fmt.Print(a ,b)
  fmt.Println()
}
