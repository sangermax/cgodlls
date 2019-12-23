#ifndef _CALL_H_
#define _CALL_H_
#include "def.h"

// C语言动态库示例，go调用

void helloworld();
void store_jpg(char* name , int len);
void only_struct(struct userinfo* u);
void multi_struct(struct userinfo* u,int size);

#endif //_CALL_H_

