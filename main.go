package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args)<2{
    fmt.Println(0)
  } else if len(os.Args)==2 {
    if os.Args[1]=="'ok'" { 
      fmt.Println(1)
    } else {
      fmt.Println(0)
    }
  } else {
    fmt.Println(len(os.Args)-1)
  }
  //os.Exit(0)
}  