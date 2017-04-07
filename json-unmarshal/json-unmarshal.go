package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Example struct {
	String      string   `json:"string"`
	Int         int      `json:"int"`
	Float       float64  `json:"float"`
	StringArray []string `json:"string-array"`
}

func main() {
	var e Example
	input := `
{
	"string" : "string"
	, "int" : 4
	, "float" : 4.5
	, "string-array" : ["string-1", "string-2"]
}
`
	err := json.Unmarshal([]byte(input), &e)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", e)
	fmt.Println()
}
