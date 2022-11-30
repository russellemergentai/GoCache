package myPackage

import (
	"fmt"
"errors"
)

func Hello() {
	fmt.Println("Hello, starting...")
}

var BIGERROR = errors.New("big error")

func ThrowStuff() error {
  return BIGERROR
}
