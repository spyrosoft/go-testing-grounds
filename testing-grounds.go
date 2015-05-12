package main

import (
	"fmt"
	"os"
	"encoding/json"
)



type dummy struct {
	tweedle_dee uint16
	tweedle_dum uint16
}

type Message struct {
	Name string `json:"example-name"`
	Body string "Foo"
}



func struct_initialization() {
	eeeeeeee := dummy{}
	fmt.Println( eeeeeeee.tweedle_dum )
}

func json_example() {
	var unmarshaling_with_tag Message
	json_input := []byte( `{"example-name":"Eeeeee!"}` )
	err := json.Unmarshal( json_input, &unmarshaling_with_tag )
	if err != nil {
		os.Exit( 1 )
	}
	fmt.Println( unmarshaling_with_tag )
}



func main () {
	json_example()
	struct_initialization()
}