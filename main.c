// C调用动态库示例

#include <stdio.h>
#include "def.h"
#include "libplugin.h"
#include "clib.h"

typedef void (*Callback)(int d);

void gocallfunc(int d) {
    printf("Go动态库传递参数至此 %d\n",d);
}

void gocallfuncstr(void *s){
    printf("go回调返回 %s\n", (char *)s);
}

void some_c_func(callback_fcn call) {
    int args = 2;
    printf("c.some_c_func() %d",args);
    int response = call(2);
    printf("c.some_c_func() %d", response);
}




// 调用 libhello.so  测试动态库
int main()
{
    // c调用c动态库的demo
    struct userinfo u = {"testproc",sizeof("testproc")};
    //only_struct(&u);

    // c 调用go的函数实现
    //DCall();
    register_callback(gocallfuncstr);
    wait_event();
    // c 调用go的回调函数
    //callOnMeGo_cgo(1);
    //some_c_func(callOnMeGo_cgo);

    //callOnMeGo(1);

    //testfunc(gocallfunc);
    //Readrsudata(gocallfunc);
    

    //go_callback_int(gocallfunc, 1);




    return 0;
}
