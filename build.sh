go build -o plugin/plugin.so -buildmode=plugin plugin.go
go build -o plugin/plugin_name.so -buildmode=plugin plugin_name.go
go build -o plugin/greet.so -buildmode=plugin greet.go
go build
