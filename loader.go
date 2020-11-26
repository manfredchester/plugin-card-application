package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

type loader struct {
	pluginDir string
	objectDir string
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
	srcPath := filepath.Join(l.objectDir, fmt.Sprintf("%s.go", name))
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
