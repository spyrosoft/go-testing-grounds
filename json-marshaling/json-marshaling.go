package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Message struct {
	Name string `json:"example-name"`
}

func main() {
	var unmarshaling_with_tag Message
	json_input := []byte( `{"example-name":"Eeeeee!"}` )
	err := json.Unmarshal( json_input, &unmarshaling_with_tag )
	if err != nil {
		os.Exit( 1 )
	}
	fmt.Println( unmarshaling_with_tag )
}