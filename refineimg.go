/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.10
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: refineimg.i

package refineimg

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
extern void _wrap_Swig_free_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_refineimg_4998b77f6f6cb248(swig_intgo arg1);
extern uintptr_t _wrap_new_ByteVector__SWIG_0_refineimg_4998b77f6f6cb248(void);
extern uintptr_t _wrap_new_ByteVector__SWIG_1_refineimg_4998b77f6f6cb248(swig_type_1 arg1);
extern swig_type_2 _wrap_ByteVector_size_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern swig_type_3 _wrap_ByteVector_capacity_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern void _wrap_ByteVector_reserve_refineimg_4998b77f6f6cb248(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_ByteVector_isEmpty_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern void _wrap_ByteVector_clear_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern void _wrap_ByteVector_add_refineimg_4998b77f6f6cb248(uintptr_t arg1, char arg2);
extern char _wrap_ByteVector_get_refineimg_4998b77f6f6cb248(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_ByteVector_set_refineimg_4998b77f6f6cb248(uintptr_t arg1, swig_intgo arg2, char arg3);
extern void _wrap_delete_ByteVector_refineimg_4998b77f6f6cb248(uintptr_t arg1);
extern uintptr_t _wrap_refine_refineimg_4998b77f6f6cb248(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_refineimg_4998b77f6f6cb248(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrByteVector uintptr

func (p SwigcptrByteVector) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrByteVector) SwigIsByteVector() {
}

func NewByteVector__SWIG_0() (_swig_ret ByteVector) {
	var swig_r ByteVector
	swig_r = (ByteVector)(SwigcptrByteVector(C._wrap_new_ByteVector__SWIG_0_refineimg_4998b77f6f6cb248()))
	return swig_r
}

func NewByteVector__SWIG_1(arg1 int64) (_swig_ret ByteVector) {
	var swig_r ByteVector
	_swig_i_0 := arg1
	swig_r = (ByteVector)(SwigcptrByteVector(C._wrap_new_ByteVector__SWIG_1_refineimg_4998b77f6f6cb248(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewByteVector(a ...interface{}) ByteVector {
	argc := len(a)
	if argc == 0 {
		return NewByteVector__SWIG_0()
	}
	if argc == 1 {
		return NewByteVector__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrByteVector) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ByteVector_size_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrByteVector) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ByteVector_capacity_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrByteVector) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ByteVector_reserve_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrByteVector) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_ByteVector_isEmpty_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrByteVector) Clear() {
	_swig_i_0 := arg1
	C._wrap_ByteVector_clear_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrByteVector) Add(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ByteVector_add_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrByteVector) Get(arg2 int) (_swig_ret byte) {
	var swig_r byte
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (byte)(C._wrap_ByteVector_get_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrByteVector) Set(arg2 int, arg3 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_ByteVector_set_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.char(_swig_i_2))
}

func DeleteByteVector(arg1 ByteVector) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_ByteVector_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0))
}

type ByteVector interface {
	Swigcptr() uintptr
	SwigIsByteVector()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 byte)
	Get(arg2 int) (_swig_ret byte)
	Set(arg2 int, arg3 byte)
}

func Refine(arg1 ByteVector) (_swig_ret ByteVector) {
	var swig_r ByteVector
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (ByteVector)(SwigcptrByteVector(C._wrap_refine_refineimg_4998b77f6f6cb248(C.uintptr_t(_swig_i_0))))
	return swig_r
}

