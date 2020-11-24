package main

import (
	"fmt"
	"plugin"
)

type Greeter interface {
	Greet()
}

func main() {
	fmt.Println("plug start")
	p, err := plugin.Open("plugin/plugin.so")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(p)
	m, err := p.Lookup("Greeter")
	if err != nil {
		fmt.Println("err:", err)
	}
	greeter, ok := m.(Greeter)
	if !ok {
		fmt.Println("err:", err)
	}
	greeter.Greet()
	// res := m.(func(int) int)(30)
	// fmt.Println(res)
}
