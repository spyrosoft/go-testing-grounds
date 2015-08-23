package main

import (
	"fmt"
)

func main() {
	fmt.Println("Demonstrate defer stacking")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End of function")
}