package main
import (
  "fmt"
  "os"
)

func main(){
  fmt.Println("args",os.Args)

  if len(os.Args)>1{
    fmt.Println("Hello Golang",os.Args[1])
  }
  os.Exit(01)
}
