package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	demo()

	l := newLoader()
	if l != nil {
		for _, name := range l.getplugins() {
			if err := l.compileAndRun(name); err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
		}
	}
	ch := make(chan int, 0)
	<-ch
}

func newLoader() *loader {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(fmt.Errorf("%v", e))
			return
		}
	}()
	wd, err := os.Getwd()
	asset(err)
	fmt.Println("wd:", wd)

	pDir := filepath.Join(wd, "plugin")
	fmt.Println("pDir:", pDir)

	tmp, err := ioutil.TempDir("/tmp/zhxu", "")
	asset(err)
	fmt.Println("tmp:", tmp)

	return &loader{
		pluginDir: pDir,
		objectDir: tmp,
	}
}

func asset(err error) {
	if err != nil {
		panic(err)
	}
}

type Greeter interface {
	Greet()
	// Greet1s()
}

func demo() {
	fmt.Println("**************start**********************")
	fmt.Println("==========================\necho ./plugin/greet.so")
	// 3. 查找并实例化插件
	plug1, err := plugin.Open("./plugin/greet.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(plug1)
	// 4. 找到插件导出的接口实例，其实这个不是必须的
	symGreeter, err := plug1.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 5. 类型转换
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}

	// 6. 调用方法
	greeter.Greet()
	// greeter.Greet1s()

	fmt.Println("====================================================\necho ./plugin/plugin_name.so")
	p, err := plugin.Open("./plugin/plugin_name.so")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(p)
	v, err := p.Lookup("V")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(v)
	f, err := p.Lookup("F")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	*v.(*int) = 7
	f.(func())()
	// prints "Hello, number 7"

	fmt.Println("====================\n./plugin/plugin.so")
	plugs, err := plugin.Open("./plugin/plugin.so")
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	fmt.Println(plugs)
	m, err := plugs.Lookup("Demo")
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	fmt.Println(m)
	res := m.(func(int) int)(30)
	fmt.Println(res)
	fmt.Println("**************end**********************")
}
