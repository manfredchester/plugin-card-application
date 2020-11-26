package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

type Greeter interface {
	Greet()
	// Greet1s()
}

func main() {
	l := newLoader()
	if l == nil {
		return
	}
	fmt.Println(l.getplugins())
	ch := make(chan int, 0)
	for _, name := range l.getplugins() {
		if err := l.compileAndRun(name); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}
	<-ch
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

func (l *loader) compileAndRun(name string) error {
	obj, err := l.compile(name)
	if err != nil {
		return fmt.Errorf("could not compile %s: %v", name, err)
	}
	// defer os.Remove(obj)

	if err := l.run(obj); err != nil {
		return fmt.Errorf("could not run %s: %v", obj, err)
	}
	return nil
}

func (l *loader) compile(name string) (string, error) {
	f, err := ioutil.ReadFile(filepath.Join(l.pluginDir, name))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	srcPath := filepath.Join(l.objectDir, fmt.Sprintf("%d.go", rand.Int()))
	if err := ioutil.WriteFile(srcPath, f, 0666); err != nil {
		fmt.Println(err)
		return "", err
	}
	objectPath := srcPath[:len(srcPath)-3] + ".so"
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o="+objectPath, srcPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("could not compile %s: %v", name, err)
	}

	return objectPath, nil
}

func (l *loader) run(object string) error {
	p, err := plugin.Open(object)
	if err != nil {
		return fmt.Errorf("could not open %s: %v", object, err)
	}
	run, err := p.Lookup("Demo")
	if err != nil {
		return fmt.Errorf("could not find Run function: %v", err)
	}
	// runFunc, ok := run.(func() error)
	res := run.(func(int) int)(30)
	fmt.Println(res)
	// if !ok {
	// 	return fmt.Errorf("found Run but type is %T instead of func() error", run)
	// }
	// if err := runFunc(); err != nil {
	// 	return fmt.Errorf("plugin failed with error %v", err)
	// }
	return nil

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

// func (l *loader) abandon() {
// 	os.RemoveAll(l.pluginDir)
// }

func asset(err error) {
	if err != nil {
		panic(err)
	}
}
