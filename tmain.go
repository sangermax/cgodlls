package main
// Go调用动态库示例

/*
#cgo CFLAGS : -I.
#cgo LDFLAGS: -L. -lhello
#include "hello.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func addNum(in int) int {
	return in + 1
}

func showMsg(msg *C.char) int {
	fmt.Println("sow msg in go", C.GoString(msg))
	return 1
}

func main() {
	call_helloworl()
	call_store_file()
	call_only_struct()
	call_multi_struct()
}

func call_helloworl() {
	C.helloworld()
}

func call_store_file() {
	f, err := os.Open("1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 1024*10)
	if n, err := f.Read(buf); err != nil {
		panic(err)
	} else {
		cs := C.CString(string(buf[0:n]))
		defer C.free(unsafe.Pointer(cs))
		C.store_jpg(cs, C.int(n))
	}
}

func call_only_struct() {
	var userinfo C.struct_userinfo
	cs := C.CString("shadiao wangyou")
	defer C.free(unsafe.Pointer(cs))
	userinfo.username = cs
	userinfo.len = C.int(len("shadiao wangyou"))
	C.only_struct(&userinfo)
}

func call_multi_struct() {
	var userinfos []C.struct_userinfo
	for i := 0; i < 2; i++ {
		var ui C.struct_userinfo
		s := fmt.Sprintf("%s-%d", "shadiao wangyou", i)
		cs := C.CString(s)
		defer C.free(unsafe.Pointer(cs))
		ui.username = cs
		ui.len = C.int(len(s))

		userinfos = append(userinfos, ui)
	}
	C.multi_struct(&userinfos[0], C.int(len(userinfos)))
}
