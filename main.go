package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("plug start")
	p, err := plugin.Open("plugin/plugin.so")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(p)
	m, err := p.Lookup("Demo")
	if err != nil {
		fmt.Println("err:", err)
	}
	res := m.(func(int) int)(30)
	fmt.Println(res)
}
