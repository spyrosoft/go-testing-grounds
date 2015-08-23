package main

import (
	"fmt"
)

func main() {
	switch example := "example"; example {
	case "other":
		fmt.Println("Other")
	case "example":
		fmt.Println("Example")
		fallthrough
	case "twizzler":
		fmt.Println("Twizzler")
	default:
		fmt.Printf("Default")
	}
}