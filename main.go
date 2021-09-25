package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Print(0)
/*
  if len(os.Args)<2{
    fmt.Print(0)
  } else if os.Args[1]=="''" {
    fmt.Print(0)
  } else {
    fmt.Print(len(os.Args)-1)
  }*/
  os.Exit(0)
}  