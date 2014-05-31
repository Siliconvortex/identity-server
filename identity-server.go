package main

import (
  "fmt"
  "github.com/dchest/uniuri"
)

func main() {
  fmt.Printf(uniuri.NewLen(256))
  fmt.Printf("\n")
}
