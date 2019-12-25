// C调用动态库示例

#include <stdio.h>
#include "libplugin.h"


void gocallfuncstr(void *s, int len){
    printf("收到天线返回 长度:%d : %s\n", len, (char *)s);

    //.. 此处做逻辑判断

    Writefee(100,200);


}

// 调用 libhello.so  测试动态库
int main()
{
    // 注册回调
    register_callback(gocallfuncstr);

    // 初始化
    initrsu("192.168.0.123",15,1,1);

    // 开始接收数据
    wait_event();

    return 0;
}
