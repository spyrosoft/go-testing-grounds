package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

// Run `./standard-input < standard-input.go` for a demonstration

func main() {
	standard_in := get_standard_in()
	fmt.Println( standard_in )
}

func get_standard_in() string {
	var standard_in_buffer []byte
	standard_in_buffer, _ = ioutil.ReadAll( os.Stdin )
	return string( standard_in_buffer )
}