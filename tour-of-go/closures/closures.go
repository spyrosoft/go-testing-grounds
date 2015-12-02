package main

import "fmt"

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		total := first + second
		first = second
		second = total
		return total
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}