package main

import (
	"fmt"
	"os"
	"plugin"
)

type Greeter interface {
	Greet()
}

func main() {
	fmt.Println("echo ./plugin/greet.so")
	// 3. 查找并实例化插件
	plug, err := plugin.Open("./plugin/greet.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 4. 找到插件导出的接口实例，其实这个不是必须的
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 5. 类型转换
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}

	// 6. 调用方法
	greeter.Greet()

	fmt.Println("echo ./plugin/plugin_name.so")
	p, err := plugin.Open("./plugin/plugin_name.so")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"

	fmt.Println("./plugin/plugin.so")
	plugs, err := plugin.Open("./plugin/plugin.so")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(plugs)
	m, err := p.Lookup("Demo")
	if err != nil {
		fmt.Println("err:", err)
	}
	res := m.(func(int) int)(30)
	fmt.Println(res)
}
