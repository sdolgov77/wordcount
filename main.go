package main

import (
  "fmt"
  "flag"
)

func main() {
  flag.Parse()
  a := flag.Args()
  if len(a)<1 || a[0]=="''" {
    fmt.Println(0)
  } else {
    fmt.Println(len(a))
  }
}  