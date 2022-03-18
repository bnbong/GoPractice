package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // execute after the function done.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return length, uppercase // not essential, just type return
}

func main() {
	totalLength, uppercase := lenAndUpper("bnbong")
	fmt.Println(totalLength, uppercase)
}
