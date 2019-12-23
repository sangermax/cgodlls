### 编译c语言通用的动态库
go build -buildmode=c-shared -o libplugin.so plugin.go cfuncs.go

### 编译goplugin方式的动态库
go build --buildmode=plugin plugin.go

### demo c 动态库编译
gcc -shared ./test_so.c -o test_so.so

main.c  示例为c加载c、go的动态库方法

plugin.go 示例为go动态库

tmain.go 示例为go调用c动态库

### 执行
./make.sh  编译
./main  执行查看结果
