// C调用动态库示例

#include <stdio.h>
#include "cJSON.h"
#include "libplugin.h"



void gocallfuncstr(void *s, int len){
    printf("raw data:%d : %s\n", len, (char *)s);

    //.. 此处做逻辑判断

    cJSON *json;
	
	json=cJSON_Parse(s);
	if (!json) 
    {
        printf("Error before: [%s]\n",cJSON_GetErrorPtr());
    }
	else
	{
		
        cJSON *item = cJSON_GetObjectItem(json,"plate");
        printf("======= plate: %s  ======\n",item->valuestring);

        item = cJSON_GetObjectItem(json,"msgtype");
        printf("======= msg type :%d =========\n", item->valueint);


		cJSON_Delete(json);
	}


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
