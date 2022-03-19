package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // can make variable in if loop
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
