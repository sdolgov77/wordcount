package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args)<2{
    fmt.Print(0)
  } else if len(os.Args)>1 && os.Args[1]=="''" {
    fmt.Print(0)
  } else {
    fmt.Print(len(os.Args)-1)
  }
  os.Exit(0)
}  