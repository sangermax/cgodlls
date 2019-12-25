/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "main.go"


#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
typedef void (*callback)(void *);
int dealrsu_cgo(char *buffer);
static callback _cb;
static void register_callback(callback cb) {
    _cb = cb;
}
static void wait_event() {
	int i = 0;
	int c = 0;

	// 天线内容缓冲区
	char buffer[4048];

	// 不断的接收数据，回调给c客户端
	while(1) {
		memset((void *)buffer, 0 , sizeof(buffer));
		i ++;
		// 读天线，流程结束后返回c调用，代码不在此处过多写
		c = dealrsu_cgo(buffer);
		_cb(buffer);
	}
}

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


// 天线初始化

extern void initrsu(char* p0, GoInt p1, GoInt p2, GoInt p3);

extern int dealrsu(char* p0);

extern void Writefee(GoInt p0, GoInt p1);

#ifdef __cplusplus
}
#endif
