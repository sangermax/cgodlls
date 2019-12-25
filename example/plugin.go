package main

/*
#cgo CFLAGS: -I .
#include <unistd.h>
#include <stdio.h>
#include <string.h>
typedef void (*callback)(void *);
int myprintf_cgo(int i);
static callback _cb;
static void register_callback(callback cb) {
    _cb = cb;
}
static void wait_event() {
	int i = 0;
	int c = 0;
	char buffer[1024];

	// 初始化天线

	while(1) {
		memset((void *)buffer, 0 , sizeof(buffer));
		i ++;
		// 读天线，流程结束后返回c调用，代码不在此处过多写
		c = myprintf_cgo(i);
		sprintf(buffer,"test data in %d - %d", c, i);
		_cb(buffer);
	}
}
*/
import "C"
import (
	"fmt"
	"time"
	dev "cgodlls/example/device"
	//"unsafe"
	//"github.com/mattn/go-pointer"
)

type Replaydata struct {
	Error int `json:"error"`
	Errdes string `json:"errdes"`
	Obuid string `json:"obuid"`
	Cardid string `json:"cardid"`
	Plate string `json:"plate"`
}

var (
	rsu dev.Rsu
)

// 天线初始化
//export initrsu
func initrsu(rsuip *C.char, power int, channel int, waittime int) {

	rsuipstr := C.GoString(rsuip)
	rsu.Init(rsuipstr, power, channel, waittime)
	//defer C.free(unsafe.Pointer(rsuipstr))
}

//export myprintf
func myprintf(i C.int) C.int {
	time.Sleep(time.Second)
	idata := int(i)
	s := fmt.Sprintf("Go dll %d", idata)
	fmt.Println(s)

	//s := fmt.Sprintf("Go dll string %d", i)
	//return C.CString("testdata")
	return i * 100
}

//export Writefee
func Writefee(obuid int, fee int) {
	fmt.Println("扣费 ", fee, "元 ")
}

func main() {
	fmt.Println("goroute全部退出")
}
