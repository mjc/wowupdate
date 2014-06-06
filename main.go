package main

import (
  "fmt"
)

func main() {
  for _, v := range ListAddons() {
    fmt.Println(v.Name())
  }
}
