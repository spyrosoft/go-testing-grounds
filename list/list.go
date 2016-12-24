package main

import (
	"fmt"
	"container/list"
)

func addUniqueStringItemToList( string string, list *list.List ) {
	for element := list.Front(); element != nil; element = element.Next() {
		if element.Value == string { return }
	}
	list.PushBack( string )
}

func mapcar( inputList *list.List, anonymousFunction func(string) string ) *list.List {
	newList := list.New()
	for element := inputList.Front(); element != nil; element = element.Next() {
		newList.PushBack( anonymousFunction( element.Value.(string) ) )
	}
	return newList
}

func main() {
	exampleList := list.New()
	exampleList.PushBack( "first" )
	exampleList.PushBack( "second" )
	mapcar( exampleList, copyString )
}