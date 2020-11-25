package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
)

type Greeter interface {
	Greet()
	// Greet1s()
}

func main() {
	l, err := newLoader()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(l.getplugins())
	// ch:= make(chan,0)
	// for _, name := range l.getplugins() {
	// 	if err := l.compileAndRun(name); err != nil {
	// 		fmt.Fprintf(os.Stderr, "%v", err)
	// 	}
	// }
	// ch<-
	return
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

type loader struct {
	pluginDir string
	objectDir string
}

func (l *loader) compileAndRun() {
	// obj,err:=l.compile()
}

func (l *loader) compile(name string) (string, error) {
	f, err := ioutil.ReadFile(filepath.Join(l.pluginDir, name))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

}

func (l *loader) run() {

}

func newLoader() (*loader, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	pDir := filepath.Join(wd, "plugin")

	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &loader{
		pluginDir: pDir,
		objectDir: tmp,
	}, nil

}

func (l *loader) getplugins() []string {
	dir, err := os.Open(l.pluginDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var res []string
	for _, name := range names {
		if filepath.Ext(name) == ".go" {
			res = append(res, name)
		}
	}
	return res
}
