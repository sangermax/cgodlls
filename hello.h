#ifndef _CALL_H_
#define _CALL_H_
#include "def.h"

// C语言动态库示例，go调用

typedef int (*fnc)(int); // 回调函数的名称为 fnc，参数是int
typedef int (*fnc2)(char*); // 回调函数的名称为 fnc2，参数是 char *str



void helloworld();
void store_jpg(char* name , int len);
void only_struct(struct userinfo* u);
void multi_struct(struct userinfo* u,int size);
void callBack(fnc notifyMain);
void callBack2(fnc2 notifyMain);

#endif //_CALL_H_

