package main

import (
	"fmt"
	// myPackage "main/myPackage"
)

// func SumIntsOrFloats[K comparable, V float64](m map[K]V) V {
//     var s V
//     for _, v := range m {
//         s += v
//     }
//     return s
// }

func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

      
func main() {

floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

  acc := SumFloats(floats) 
  
fmt.Printf("%f", acc)
  
}