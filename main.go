package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

// not exist at Go:
// constructor()
// __init__

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name:"nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}
