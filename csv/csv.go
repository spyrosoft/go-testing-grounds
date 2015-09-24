package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func openFile( fileName string ) *os.File {
	openedFile, error := os.Open( fileName )
	exitOnError( error )
	return openedFile
}

func printOutCsvFile( csvFile *os.File ) {
	csvReader := csv.NewReader( csvFile )
	csvContents, error := csvReader.ReadAll()
	exitOnError( error )
	for _, row := range csvContents {
		fmt.Println( "Fruit: " + row[ 0 ] + " | Shape : " + row[ 1 ] + " | Color: " + row[ 2 ] )
	}
}

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	csvFile := openFile( "example.csv" )
	printOutCsvFile( csvFile )
}