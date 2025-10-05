package main

import "fmt"

func double(num int) int {
  return num * 3
}

func main() {
  num := 3
  fmt.Println(double(num))
}
