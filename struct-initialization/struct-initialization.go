package main

import (
	"fmt"
)

type Example struct {
	tweedle_dee uint16
	tweedle_dum uint16
}

func main() {
	example := Example{}
	fmt.Println( example.tweedle_dum )
	fmt.Println( example )
}