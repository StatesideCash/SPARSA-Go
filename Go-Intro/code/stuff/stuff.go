package main

import "fmt"

func main() {
  s1 := []byte("01234")
  s2 := []byte("56789")
  s3 := []byte("01234")
  data := append(s1, append(s2, s3...)...)
   fmt.Printf("%s\n", data)
}
