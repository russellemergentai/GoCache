package main

import (
	"fmt"
	// myPackage "main/myPackage"
)

type Human interface {
  hobby() string
}

type Man struct {
  name string
}

type Woman struct {
  name string
}

// implement the interface fir each struct
func (m Man) hobby () string {
 return m.name
}

func (m Woman) hobby () string {
  return m.name
}
      
func main() {
  m := new (Man)
  m.name = "chad"
  w := new (Woman)
  w.name = "karen"

  fmt.Println(m.hobby())
  fmt.Println(w.hobby())

  // do polymorphism via interface

  people := []Human{m, w}

  for _, item := range people {
    fmt.Printf("%s \n", item.hobby())
  }

}