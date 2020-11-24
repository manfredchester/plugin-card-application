package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("plug start")
	p, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
