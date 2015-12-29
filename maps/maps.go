package main

import (
	"fmt"
)

func main() {
	var testMap = make( map[string]uint16 )
	testMap[ "a" ] = 4
	fmt.Println( mapKeyExists( testMap, "a" ) )
	fmt.Println( mapKeyExists( testMap, "b" ) )
}

func mapKeyExists( testMap map[string]uint16, testKey string ) bool {
	_, keyExists := testMap[ testKey ]
	return keyExists
}