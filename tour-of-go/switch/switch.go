package main

import (
	"fmt"
)

func fallthroughExample() {
	switch example := "example"; example {
	case "other":
		fmt.Println("Other")
	case "example":
		fmt.Println("Example")
		//Notice the fallthrough statement is opposite of C (Yay!!)
		fallthrough
	case "twizzler":
		fmt.Println("Twizzler")
	default:
		fmt.Printf("Default")
	}
}

func noConditionExample() {
	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		//Referring to the Yahoo Answers meme
		fmt.Println( "Good Mroing." )
	case currentTime.Hour() < 17:
		//Futurama
		fmt.Println( "Good News, Everyone!" )
	default:
		fmt.Println( "Good Evening." )
	}
}

func main() {
	fallthroughExample()
	noConditionExample()
}