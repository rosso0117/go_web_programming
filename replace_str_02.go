package main

import "fmt"

func main() {
  s := "hello"
  // byteのスライスを作成
  c := []byte(s)
  c[3] = 'c'
  s2 := string(c)
  fmt.Printf("%s\n", s2)
}
