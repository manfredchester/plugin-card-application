package main

type Greeter string

func (g Greeter) Greet() {
	println("hello world")
}
