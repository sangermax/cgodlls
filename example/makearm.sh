#!bin/sh


#freescale 提供的交叉编译库是32位的，注意安装后安装apt-get install lib32ncurses5 lib32z1 即可执行
echo "build arm so"
CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-fsl-linux-gnueabi-gcc go build -buildmode=c-shared -o libplugin.so main.go bridge.go

echo "build main"
#在飞凌的板子上运行，需装飞凌的交叉编译环境
arm-fsl-linux-gnueabi-gcc main.c cJSON.c -L. -lplugin -lm  -o main
