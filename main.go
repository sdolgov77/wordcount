package main

import (
  "os"
  "fmt"
)

func main() {
  fmt.Println(len(os.Args)-1)
}  