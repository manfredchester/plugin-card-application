package main

import (
	"fmt"
	"os"
	"plugin"
)

<<<<<<< HEAD
type Greeters interface {
=======
type Greeter interface {
>>>>>>> 0476bd76ddca860cd592ed3dd73f204e3295a147
	Greet()
}

func main() {
	// 3. 查找并实例化插件
	plug, err := plugin.Open("./greet.so")
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
	var greeter Greeters
	greeter, ok := symGreeter.(Greeters)
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}

	// 6. 调用方法
	greeter.Greet()

	p, err := plugin.Open("./plugin_name.so")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
<<<<<<< HEAD
	v, err := p.Lookup("V")
=======
	m, err := p.Lookup("Greeter")
>>>>>>> 0476bd76ddca860cd592ed3dd73f204e3295a147
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"

	// fmt.Println("plug start")
	// p, err := plugin.Open("plugin/plugin.so")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	// fmt.Println(p)
	// m, err := p.Lookup("Demo")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
=======
	greeter, ok := m.(Greeter)
	if !ok {
		fmt.Println("err:", err)
	}
	greeter.Greet()
>>>>>>> 0476bd76ddca860cd592ed3dd73f204e3295a147
	// res := m.(func(int) int)(30)
	// fmt.Println(res)
}
