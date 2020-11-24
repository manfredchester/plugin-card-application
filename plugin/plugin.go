package main

import "fmt"

func Demo(baseNum int) int {
	return baseNum + 100
}

type Greeter string

//
func (g *Greeter) Greet() {
	fmt.Println("Greet!")
}
