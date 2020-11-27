# plugin-card-application


## go-plugin是什么

1. plugin插件是包含可导出可访问的function和变量的，且是由main package编译之后的文件。注意这里main并不是业务服务中的main包而是plugin中有属于自己独立的main package。

2. plugin插件使我们可以使用动态加载链接库构建松散耦合的模块化的程序，可以在运行时动态加载和绑定。

3. 官方文档：https://golang.google.cn/pkg/plugin/#pkg-overview

## go-plugin生命周期

> plugin插件被打开plugin.Open(*.so)，插件的init()才被执行。
插件只被初始化一次，不能被关闭。

1. main.go的init函数执行
2. 开始执行main.go main函数
3. 开始执行plugin.Open("***.so")打开插件
4. 插件开始执行内部的init函数

## go-plugin应用场景

1. 通过plugin我们可以很方便的对于不同功能加载相应的模块并调用相关的模块;
2. 针对不同语言(英文,汉语,德语……)加载不同的语言so文件,进行不同的输出;
3. 编译出的文件给不同的编程语言用(如：c/java/python/lua等).
4. 需要加密的核心算法,核心业务逻辑可以可以编译成plugin插件
函数集动态加载

## go-plugin代码示例
        
编写plugin插件要点
1. 包名称必须为main
2. 没有main函数
3. 必须有可以导出可访问的变量或方法

编写完成之后编译plugin
    
使用加载plugin插件基本流程
1. 加载编译好的插件 plugin.Open("./plugin_doctor.so") (*.so文件路径相对与可执行文件的路径,可以是绝对路径)
2. 寻找插件可到变量 plug.Lookup("Greeter")
3. TypeAssert: Symbol(interface{}) 转换成API的接口类型
4. 执行API interface的方法
    
## Go-plugin局限和不足
1. go-plugin迫使插件实现与主应用程序产生了高度耦合，如果插件的作者对主应用程序没有控制权，维护的开销很更高
2. 由于插件提供的代码和主代码在相同的进程空间中运行，使得其插件实现和主应用程序都必须使用完成相同的go工具链版本构建
3. 目前支持linux和mac版本操作系统

## 总结
1. Go plugin包提供了一个简单的函数集动态加载,可以帮助开发人员编写可扩展的代码
2. Go插件是使用go build -buildmode = plugin构建标志编译
3. Go插件包中的导出函数和公开变量,可以使用插件包在运行时查找并绑定调用