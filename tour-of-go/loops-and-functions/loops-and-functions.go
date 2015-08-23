package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 4
	var i int = 0
	for ; z != math.Sqrt(x); i++ {
		z = z - (((z * z) - x) / (2 * z))
	}
	fmt.Println("Iterations:", i)
	return z
}

func approximateSqrt() {
	fmt.Println(Sqrt(2))
}

func loopOverRange() {
	//Works for slices or hashes
	powersOfTwo := make([]int, 10)
	for i := range powersOfTwo {
		powersOfTwo[i] = 1 << uint(i)
	}
	for _, value := range powersOfTwo {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	fmt.Println("Approximate Square Root using Newton's technique:")
	approximateSqrt()
	fmt.Println()
	fmt.Println("Loop over a range:")
	loopOverRange()
}
