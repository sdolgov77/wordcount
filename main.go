package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  if len(os.Args)<2{
    fmt.Println(0)
  } else if len(os.Args)==2 {
    if os.Args[1]=="''" { 
      fmt.Println(0)
    } else {
      strs := strings.Split(os.Args[1], " ")
      fmt.Println(len(strs))
    }
  } else {
    fmt.Println(len(os.Args)-1)
  }
  os.Exit(0)
}  