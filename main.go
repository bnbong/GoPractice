package main

import (
	"fmt"

	mydict "github.com/bnbong/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary["first"])
	// fmt.Println(dictionary.Search("second"))
}
