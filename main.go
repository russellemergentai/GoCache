package main

import (
	"fmt"
	"math"
	"math/rand"
  "errors"
  myPackage "main/myPackage"
)

var ErrMonteCarlo = errors.New("calc failure")

type Result struct {
  res float64
  err error
}

var results = make(chan Result, 1)

func calculate() {
	acc := 0
	n := 10000000

	for i := 0; i < n; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		xx := math.Pow(x, 2)
		yy := math.Pow(y, 2)
		r := (xx + yy)
		if r < 1.0 {
			acc++
		}
	}

	output := float64(acc) * 4 / float64(n)
  results <- Result{output, ErrMonteCarlo}
}

func main() {
  myPackage.Hello()
	go calculate()

  r := <- results 
  fmt.Println(r.res) 
  fmt.Println(r.err)
}