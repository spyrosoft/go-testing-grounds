package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 4
	var i int = 0
	for ; z != math.Sqrt(x); i++ {
		z = z - ( ((z * z) - x) / (2 * z) )
	}
	fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}