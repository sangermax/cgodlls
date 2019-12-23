#/bin/bash

./clean.sh

echo "golib"
# cfuncs.go 实现中继,注意不能少
go build -buildmode=c-shared -o libplugin.so plugin.go cfuncs.go


rm -rf ./*.o 
echo "lib"
gcc -fPIC -shared hello.c -o libhello.so

# 同时加载两个动态库
gcc main.c -L. -lplugin -lhello  -o main

