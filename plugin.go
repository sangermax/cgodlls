package main


/*
#cgo CFLAGS: -I .
#include "clib.h"
int callOnMeGo_cgo(int in); // Forward declaration.
typedef void (*callback)(void *);

static callback _cb;
static void register_callback(callback cb) {
    _cb = cb;
}
static void wait_event() {
    _cb("test data in go");
}

*/
import "C"
import (
	"fmt"
	"time"
	//"unsafe"
	//"github.com/mattn/go-pointer"
)

/*  导入go中，go函数不能直接访问指针对象，需要pointer转换
static void _register_callback(void *user_data) {
  register_callback(cb_proxy, user_data);
}

void cb_proxy(void *v);
*/

//type Callback struct {
//    Func     func(string)
//    UserData string
//}



//export cb_proxy
//func cb_proxy(v unsafe.Pointer) {
//    cb := pointer.Restore(v).(*Callback)
//    cb.Func(cb.UserData)
//}


//export DCall
func DCall() {
	fmt.Println("plugin.so Go编译的DLL被调用")
	//C._register_callback(pointer.Save(&Callback{
	//	Func: my_callback,
	//	UserData:"my-callback",
	//}))
	//C.wait_event()
}

//export callOnMeGo
func callOnMeGo(in int) int {
	fmt.Println("调用Go的回调函数.. 实现DLL动态库函数的入口功能")
	for i:= 0 ; i < 3; i ++ {
		fmt.Println("x - ", i)
		time.Sleep(time.Second)
	}
	return in + 1
}

//export DCallWithParam
func DCallWithParam(msg string, msg2 string) {
	fmt.Println("参数内容为:", msg, "+", msg2)
}

func main() {

	fmt.Println("goroute全部退出")

}
