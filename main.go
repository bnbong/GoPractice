package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // == for each or for in
		total += number
	}
	return total
	// for i:=0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
