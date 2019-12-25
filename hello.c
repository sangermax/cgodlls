#include "hello.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>




// 指针和int 模拟存文件
void store_jpg(char* name ,int len)
{
    FILE *fp = NULL;
    fp = fopen( "file.jpg" , "wb+" );
    fwrite(name, len , 1, fp );
    fclose(fp); 
}

//only_struct  调用结构体指针
void only_struct(struct userinfo* u)
{
    char* name =u->username;
    int len = u->len;
    printf("call_only_struct :%s %d\n",name,len);
}

//multi_struct 调用结构体指针，模拟结构体数组
void multi_struct(struct userinfo* u,int size)
{
    int i = 0;
    printf("call_multi_struct\n");
    while (i < size)
    {
        char* name = u[i].username;
        int len = u[i].len;
        printf("username:%s, len:%d, size:%d\n",name,len,size);
        i++;
    }    
}


void callBack(fnc notifyMain)
{
	int r= notifyMain(4);
    sleep(1);
    printf("callback C语言中回调 int %d\n",r);
}
void callBack2(fnc2 notifyMain)
{
	char* msg="msg from c";
	int r= notifyMain(msg);
	printf("callback2 C语言回调char * %d %s\n",r,msg);

}