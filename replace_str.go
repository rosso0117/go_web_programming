package main

import "fmt"

func main() {
  s := "hello"
  // helloのindex1以降のスライスと結合
  s = "c" + s[1:]
  fmt.Printf("%s\n", s)
}
