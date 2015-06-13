package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat( "file-exists.go" )
	if err == nil {
		fmt.Println( "File exists!" )
	} else {
		fmt.Println( "File does not exist." )
	}
}