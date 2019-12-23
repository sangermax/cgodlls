package main


/*
#cgo CFLAGS: -I .
#include "clib.h"
int callOnMeGo_cgo(int in); // Forward declaration.
*/
import "C"
import (
	"fmt"
	"time"
)


//export callOnMeGo
func callOnMeGo(in int) int {
	fmt.Println("调用Go的回调函数.. 实现DLL动态库函数的入口功能")
	for i:= 0 ; i < 3; i ++ {
		fmt.Println("x - ", i)
		time.Sleep(time.Second)
	}
	return in + 1
}


//export DCall
func DCall() {
	fmt.Println("plugin.so Go编译的DLL被调用")
}

//export DCallWithParam
func DCallWithParam(msg string, msg2 string) {
	fmt.Println("参数内容为:", msg, "+", msg2)
}

func main() {

	fmt.Println("goroute全部退出")

}
