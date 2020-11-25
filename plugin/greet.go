package main

type greeter string

func (g greeter) Greet() {
	println("hello world")
}

var Greeter greeter
