package main

/*
#cgo CFLAGS: -I .
#include "def.h"
static callback _cb;
static void register_callback(callback cb) {
    _cb = cb;
}
static void wait_event() {
	int i = 0;
	int c = 0;
	while(1) {
		i ++;
		myprintf(i);
		_cb("test data in go");
	}
}
*/
import "C"
import (
	"fmt"
	"time"
	//"unsafe"
	//"github.com/mattn/go-pointer"
)

type Replaydata struct {
}

//export myprintf
func myprintf(i C.int) C.int {
	idata := int(i)
	s := fmt.Sprintf("Go dll %d", idata)
	fmt.Println(s)
	time.Sleep(time.Second)
	//s := fmt.Sprintf("Go dll string %d", i)
	//return C.CString("testdata")
	return 100
}

func main() {
	fmt.Println("goroute全部退出")
}
