package main

import "fmt"

func main() {
	var ui1 uint8 = 240
	var ui2 uint8 = 100
	ui3 := ui1 + ui2
  fmt.Printf("UI:3 %v!\n", ui3)

  var b1 bool
  var b2 bool = true
  b3 := 10 > 5
  b4 := (b1 || b2) && b3
}