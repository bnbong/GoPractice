package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge{
	case 10:
		return true
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
