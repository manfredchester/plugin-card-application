go build -o plugin/plugin.so -buildmode=plugin plugin/plugin.go
go build -o plugin/plugin_name.so -buildmode=plugin plugin/plugin_name.go
go build -o plugin/greet.so -buildmode=plugin plugin/greet.go
go build
