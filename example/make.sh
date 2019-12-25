#/bin/bash

./clean.sh

echo "编译libplugin"

go build -buildmode=c-shared -o libplugin.so plugin.go bridge.go

echo "编译c语言示例"
gcc main.c -L. -lplugin  -o main

 export LD_LIBRARY_PATH=./
 