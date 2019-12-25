// C调用动态库示例

#include <stdio.h>
#include "libplugin.h"

typedef void (*Callback)(int d);


void gocallfuncstr(void *s){
    printf("go回调返回 %s\n", (char *)s);
}

// 调用 libhello.so  测试动态库
int main()
{
    // 注册回调
    register_callback(gocallfuncstr);
    // 开始接收数据
    wait_event();

    return 0;
}
