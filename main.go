package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args)<2{
    fmt.Println(0)
  } else if len(os.Args)==2 {
      fmt.Println(3)
  } else {
    fmt.Println(len(os.Args)-1)
  }
  os.Exit(0)
}  