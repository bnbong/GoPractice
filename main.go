package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "12"} //map[key]value
	for key, _ := range nico {
		fmt.Println(key)
	}
}
