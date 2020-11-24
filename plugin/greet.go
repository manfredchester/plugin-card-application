package main

type greeting string

func (g greeting) Greet() {
	println("hello world")
}

var Greeter greeting
