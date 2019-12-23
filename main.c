// C调用动态库示例

#include <stdio.h>
#include "def.h"
#include "libplugin.h"
#include "clib.h"


void some_c_func(callback_fcn callback) {
    int args = 2;
    printf("c.some_c_func() %d",args);
    int response = callback(2);
    printf("c.some_c_func() %d", response);
}



// 调用 libhello.so  测试动态库
int main()
{
    // c调用c动态库的demo
    helloworld();
    struct userinfo u = {"testproc",sizeof("testproc")};
    only_struct(&u);

    // c 调用go的函数实现
    DCall();
    // c 调用go的回调函数
    callOnMeGo_cgo(1);
    some_c_func(callOnMeGo_cgo);
    return 0;
}
