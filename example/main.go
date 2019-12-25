package main

/*
#cgo CFLAGS: -I .
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
typedef void (*callback)(void *,int);
int dealrsu_cgo(char *buffer);
static callback _cb;
static void register_callback(callback cb) {
    _cb = cb;
}
static void wait_event() {
	int ilen = 0;
	// 天线内容缓冲区
	char buffer[4048];

	// 不断的接收数据，回调给c客户端
	while(1) {
		memset((void *)buffer, 0 , sizeof(buffer));
		// 读天线，流程结束后返回c调用，代码不在此处过多写
		ilen = dealrsu_cgo(buffer);
		_cb(buffer ,ilen);
	}
}
*/
import "C"
import (
	"fmt"
	"time"
	"encoding/json"
	dev "cgodlls/example/device"
	"unsafe"
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
}

//export dealrsu
func dealrsu(buffer* C.char) int {

	var replaydata Replaydata
	replaydata.Plate = "苏A12345"
	replaydata.Obuid = "123456"
	replaydata.Cardid = "测试卡320100000000011223456789"

	jsonarr,_ := json.Marshal(replaydata)
	strjson := string(jsonarr)

	// 此处只做个模拟，实际实现go语言读取天线的逻辑,拷贝结果数据到函数里面


	var cmsg *C.char = C.CString(strjson)
	lenstr := len(strjson)
	C.memcpy(unsafe.Pointer(buffer), unsafe.Pointer(cmsg),C.size_t(lenstr))
	defer C.free(unsafe.Pointer(cmsg))

	

	time.Sleep(time.Second)
	return lenstr
}

//export Writefee
func Writefee(obuid int, fee int) {
	fmt.Println("扣费 ", fee, "元 ")
}

func main() {
	fmt.Println("goroute全部退出")
}
