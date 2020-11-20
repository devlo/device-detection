/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: src\pattern\51Degrees.i

package FiftyOneDegreesPatternV3


/*
#cgo LDFLAGS: -lpthread
#include <time.h>
#include "../threading.c"
#include "../cityhash/city.c"
*/
import "C"

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef long long swig_type_7;
typedef long long swig_type_8;
typedef long long swig_type_9;
typedef long long swig_type_10;
typedef _gostring_ swig_type_11;
typedef _gostring_ swig_type_12;
typedef _gostring_ swig_type_13;
typedef _gostring_ swig_type_14;
typedef _gostring_ swig_type_15;
typedef _gostring_ swig_type_16;
typedef _gostring_ swig_type_17;
typedef _gostring_ swig_type_18;
typedef _gostring_ swig_type_19;
typedef _gostring_ swig_type_20;
typedef _gostring_ swig_type_21;
typedef _gostring_ swig_type_22;
typedef _gostring_ swig_type_23;
typedef _gostring_ swig_type_24;
typedef _gostring_ swig_type_25;
typedef _gostring_ swig_type_26;
typedef _gostring_ swig_type_27;
typedef _gostring_ swig_type_28;
typedef _gostring_ swig_type_29;
typedef _gostring_ swig_type_30;
typedef _gostring_ swig_type_31;
typedef _gostring_ swig_type_32;
typedef _gostring_ swig_type_33;
typedef _gostring_ swig_type_34;
typedef _gostring_ swig_type_35;
typedef _gostring_ swig_type_36;
typedef _gostring_ swig_type_37;
typedef _gostring_ swig_type_38;
typedef _gostring_ swig_type_39;
typedef _gostring_ swig_type_40;
typedef _gostring_ swig_type_41;
typedef _gostring_ swig_type_42;
typedef _gostring_ swig_type_43;
typedef _gostring_ swig_type_44;
extern void _wrap_Swig_free_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_intgo arg1);
extern uintptr_t _wrap_new_MapStringString__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(void);
extern uintptr_t _wrap_new_MapStringString__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_MapStringString_size_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern _Bool _wrap_MapStringString_empty_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_MapStringString_clear_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_1 _wrap_MapStringString_get_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_2 arg2);
extern void _wrap_MapStringString_set_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_3 arg2, swig_type_4 arg3);
extern void _wrap_MapStringString_del_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_5 arg2);
extern _Bool _wrap_MapStringString_has_key_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_6 arg2);
extern void _wrap_delete_MapStringString_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_new_VectorString__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(void);
extern uintptr_t _wrap_new_VectorString__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_7 arg1);
extern swig_type_8 _wrap_VectorString_size_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_9 _wrap_VectorString_capacity_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_VectorString_reserve_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_10 arg2);
extern _Bool _wrap_VectorString_isEmpty_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_VectorString_clear_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_VectorString_add_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_11 arg2);
extern swig_type_12 _wrap_VectorString_get_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_VectorString_set_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2, swig_type_13 arg3);
extern void _wrap_delete_VectorString_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_delete_Match_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_Match_getValues__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_14 arg2);
extern uintptr_t _wrap_Match_getValues__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_voidp arg2);
extern uintptr_t _wrap_Match_getValues__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2);
extern swig_type_15 _wrap_Match_getValue__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_16 arg2);
extern swig_type_17 _wrap_Match_getValue__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_voidp arg2);
extern swig_type_18 _wrap_Match_getValue__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2);
extern swig_type_19 _wrap_Match_getDeviceId_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Match_getRank_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Match_getDifference_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Match_getMethod_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_20 _wrap_Match_getUserAgent_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_delete_Profiles_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_new_Profiles_FiftyOneDegreesPatternV3_c8b75028699b5788(void);
extern swig_intgo _wrap_Profiles_getCount_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Profiles_getProfileIndex_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_Profiles_getProfileId_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_new_Provider__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_21 arg1);
extern uintptr_t _wrap_new_Provider__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_22 arg1, swig_type_23 arg2);
extern uintptr_t _wrap_new_Provider__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_24 arg1, swig_type_25 arg2, swig_intgo arg3, swig_intgo arg4);
extern uintptr_t _wrap_new_Provider__SWIG_3_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_26 arg1, uintptr_t arg2);
extern uintptr_t _wrap_new_Provider__SWIG_4_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_27 arg1, uintptr_t arg2, swig_intgo arg3, swig_intgo arg4);
extern uintptr_t _wrap_new_Provider__SWIG_5_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_28 arg1, swig_intgo arg2, swig_intgo arg3);
extern void _wrap_delete_Provider_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_Provider_getHttpHeaders_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_Provider_getAvailableProperties_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_29 _wrap_Provider_getDataSetName_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_30 _wrap_Provider_getDataSetFormat_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_31 _wrap_Provider_getDataSetPublishedDate_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_type_32 _wrap_Provider_getDataSetNextUpdateDate_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Provider_getDataSetSignatureCount_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Provider_getDataSetDeviceCombinations_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_Provider_getMatch__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_33 arg2);
extern uintptr_t _wrap_Provider_getMatch__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, uintptr_t arg2);
extern swig_type_34 _wrap_Provider_getMatchJson__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_35 arg2);
extern swig_type_36 _wrap_Provider_getMatchJson__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_Provider_getMatchForDeviceId_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_37 arg2);
extern uintptr_t _wrap_Provider_findProfiles__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_38 arg2, swig_type_39 arg3);
extern uintptr_t _wrap_Provider_findProfiles__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_40 arg2, swig_type_41 arg3, uintptr_t arg4);
extern void _wrap_Provider_reloadFromFile_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern void _wrap_Provider_reloadFromMemory_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1, swig_type_42 arg2, swig_intgo arg3);
extern swig_intgo _wrap_Provider_getCacheHits_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Provider_getCacheMisses_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern swig_intgo _wrap_Provider_getCacheMaxIterations_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern _Bool _wrap_Provider_getIsThreadSafe_FiftyOneDegreesPatternV3_c8b75028699b5788(uintptr_t arg1);
extern uintptr_t _wrap_new_Provider__SWIG_6_FiftyOneDegreesPatternV3_c8b75028699b5788(swig_type_43 arg1, swig_type_44 arg2, swig_intgo arg3, swig_intgo arg4, _Bool arg5);
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


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_FiftyOneDegreesPatternV3_c8b75028699b5788(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrMapStringString uintptr

func (p SwigcptrMapStringString) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMapStringString) SwigIsMapStringString() {
}

func NewMapStringString__SWIG_0() (_swig_ret MapStringString) {
	var swig_r MapStringString
	swig_r = (MapStringString)(SwigcptrMapStringString(C._wrap_new_MapStringString__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788()))
	return swig_r
}

func NewMapStringString__SWIG_1(arg1 MapStringString) (_swig_ret MapStringString) {
	var swig_r MapStringString
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (MapStringString)(SwigcptrMapStringString(C._wrap_new_MapStringString__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewMapStringString(a ...interface{}) MapStringString {
	argc := len(a)
	if argc == 0 {
		return NewMapStringString__SWIG_0()
	}
	if argc == 1 {
		return NewMapStringString__SWIG_1(a[0].(MapStringString))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrMapStringString) Size() (_swig_ret uint) {
	var swig_r uint
	_swig_i_0 := arg1
	swig_r = (uint)(C._wrap_MapStringString_size_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMapStringString) Empty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_MapStringString_empty_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMapStringString) Clear() {
	_swig_i_0 := arg1
	C._wrap_MapStringString_clear_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrMapStringString) Get(arg2 string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_MapStringString_get_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrMapStringString) Set(arg2 string, arg3 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_MapStringString_set_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_2)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
}

func (arg1 SwigcptrMapStringString) Del(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_MapStringString_del_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrMapStringString) Has_key(arg2 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_MapStringString_has_key_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func DeleteMapStringString(arg1 MapStringString) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_MapStringString_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

type MapStringString interface {
	Swigcptr() uintptr
	SwigIsMapStringString()
	Size() (_swig_ret uint)
	Empty() (_swig_ret bool)
	Clear()
	Get(arg2 string) (_swig_ret string)
	Set(arg2 string, arg3 string)
	Del(arg2 string)
	Has_key(arg2 string) (_swig_ret bool)
}

type SwigcptrVectorString uintptr

func (p SwigcptrVectorString) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorString) SwigIsVectorString() {
}

func NewVectorString__SWIG_0() (_swig_ret VectorString) {
	var swig_r VectorString
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_new_VectorString__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788()))
	return swig_r
}

func NewVectorString__SWIG_1(arg1 int64) (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_new_VectorString__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.swig_type_7(_swig_i_0))))
	return swig_r
}

func NewVectorString(a ...interface{}) VectorString {
	argc := len(a)
	if argc == 0 {
		return NewVectorString__SWIG_0()
	}
	if argc == 1 {
		return NewVectorString__SWIG_1(a[0].(int64))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVectorString) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VectorString_size_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorString) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VectorString_capacity_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorString) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VectorString_reserve_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_type_10(_swig_i_1))
}

func (arg1 SwigcptrVectorString) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_VectorString_isEmpty_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorString) Clear() {
	_swig_i_0 := arg1
	C._wrap_VectorString_clear_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVectorString) Add(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VectorString_add_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_11)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrVectorString) Get(arg2 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_VectorString_get_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrVectorString) Set(arg2 int, arg3 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_VectorString_set_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), *(*C.swig_type_13)(unsafe.Pointer(&_swig_i_2)))
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
}

func DeleteVectorString(arg1 VectorString) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_VectorString_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

type VectorString interface {
	Swigcptr() uintptr
	SwigIsVectorString()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 string)
	Get(arg2 int) (_swig_ret string)
	Set(arg2 int, arg3 string)
}

type SwigcptrMatch uintptr

func (p SwigcptrMatch) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrMatch) SwigIsMatch() {
}

func DeleteMatch(arg1 Match) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Match_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrMatch) GetValues__SWIG_0(arg2 string) (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_Match_getValues__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_14)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrMatch) GetValues__SWIG_1(arg2 *string) (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_Match_getValues__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrMatch) GetValues__SWIG_2(arg2 int) (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_Match_getValues__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (p SwigcptrMatch) GetValues(a ...interface{}) VectorString {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(*string); !ok {
			goto check_1
		}
		return p.GetValues__SWIG_1(a[0].(*string))
	}
check_1:
	if argc == 1 {
		if _, ok := a[0].(int); !ok {
			goto check_2
		}
		return p.GetValues__SWIG_2(a[0].(int))
	}
check_2:
	if argc == 1 {
		return p.GetValues__SWIG_0(a[0].(string))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrMatch) GetValue__SWIG_0(arg2 string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_Match_getValue__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_16)(unsafe.Pointer(&_swig_i_1)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrMatch) GetValue__SWIG_1(arg2 *string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_Match_getValue__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrMatch) GetValue__SWIG_2(arg2 int) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_Match_getValue__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (p SwigcptrMatch) GetValue(a ...interface{}) string {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(*string); !ok {
			goto check_1
		}
		return p.GetValue__SWIG_1(a[0].(*string))
	}
check_1:
	if argc == 1 {
		if _, ok := a[0].(int); !ok {
			goto check_2
		}
		return p.GetValue__SWIG_2(a[0].(int))
	}
check_2:
	if argc == 1 {
		return p.GetValue__SWIG_0(a[0].(string))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrMatch) GetDeviceId() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Match_getDeviceId_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrMatch) GetRank() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Match_getRank_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMatch) GetDifference() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Match_getDifference_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMatch) GetMethod() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Match_getMethod_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrMatch) GetUserAgent() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Match_getUserAgent_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

type Match interface {
	Swigcptr() uintptr
	SwigIsMatch()
	GetValues(a ...interface{}) VectorString
	GetValue(a ...interface{}) string
	GetDeviceId() (_swig_ret string)
	GetRank() (_swig_ret int)
	GetDifference() (_swig_ret int)
	GetMethod() (_swig_ret int)
	GetUserAgent() (_swig_ret string)
}

type SwigcptrProfiles uintptr

func (p SwigcptrProfiles) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrProfiles) SwigIsProfiles() {
}

func DeleteProfiles(arg1 Profiles) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Profiles_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func NewProfiles() (_swig_ret Profiles) {
	var swig_r Profiles
	swig_r = (Profiles)(SwigcptrProfiles(C._wrap_new_Profiles_FiftyOneDegreesPatternV3_c8b75028699b5788()))
	return swig_r
}

func (arg1 SwigcptrProfiles) GetCount() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Profiles_getCount_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProfiles) GetProfileIndex(arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_Profiles_getProfileIndex_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrProfiles) GetProfileId(arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_Profiles_getProfileId_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

type Profiles interface {
	Swigcptr() uintptr
	SwigIsProfiles()
	GetCount() (_swig_ret int)
	GetProfileIndex(arg2 int) (_swig_ret int)
	GetProfileId(arg2 int) (_swig_ret int)
}

type SwigcptrProvider uintptr

func (p SwigcptrProvider) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrProvider) SwigIsProvider() {
}

func NewProvider__SWIG_0(arg1 string) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_21)(unsafe.Pointer(&_swig_i_0)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func NewProvider__SWIG_1(arg1 string, arg2 string) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_22)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_23)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewProvider__SWIG_2(arg1 string, arg2 string, arg3 int, arg4 int) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_2_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_24)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_25)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewProvider__SWIG_3(arg1 string, arg2 VectorString) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_3_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_26)(unsafe.Pointer(&_swig_i_0)), C.uintptr_t(_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func NewProvider__SWIG_4(arg1 string, arg2 VectorString, arg3 int, arg4 int) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_4_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_27)(unsafe.Pointer(&_swig_i_0)), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func NewProvider__SWIG_5(arg1 string, arg2 int, arg3 int) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_5_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_28)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func DeleteProvider(arg1 Provider) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Provider_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrProvider) GetHttpHeaders() (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_Provider_getHttpHeaders_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrProvider) GetAvailableProperties() (_swig_ret VectorString) {
	var swig_r VectorString
	_swig_i_0 := arg1
	swig_r = (VectorString)(SwigcptrVectorString(C._wrap_Provider_getAvailableProperties_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrProvider) GetDataSetName() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Provider_getDataSetName_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrProvider) GetDataSetFormat() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Provider_getDataSetFormat_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrProvider) GetDataSetPublishedDate() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Provider_getDataSetPublishedDate_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrProvider) GetDataSetNextUpdateDate() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Provider_getDataSetNextUpdateDate_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrProvider) GetDataSetSignatureCount() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Provider_getDataSetSignatureCount_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProvider) GetDataSetDeviceCombinations() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Provider_getDataSetDeviceCombinations_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProvider) GetMatch__SWIG_0(arg2 string) (_swig_ret Match) {
	var swig_r Match
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Match)(SwigcptrMatch(C._wrap_Provider_getMatch__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_33)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrProvider) GetMatch__SWIG_1(arg2 MapStringString) (_swig_ret Match) {
	var swig_r Match
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Match)(SwigcptrMatch(C._wrap_Provider_getMatch__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

func (p SwigcptrProvider) GetMatch(a ...interface{}) Match {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(string); !ok {
			goto check_1
		}
		return p.GetMatch__SWIG_0(a[0].(string))
	}
check_1:
	if argc == 1 {
		return p.GetMatch__SWIG_1(a[0].(MapStringString))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrProvider) GetMatchJson__SWIG_0(arg2 string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r_p := C._wrap_Provider_getMatchJson__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_35)(unsafe.Pointer(&_swig_i_1)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrProvider) GetMatchJson__SWIG_1(arg2 MapStringString) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r_p := C._wrap_Provider_getMatchJson__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (p SwigcptrProvider) GetMatchJson(a ...interface{}) string {
	argc := len(a)
	if argc == 1 {
		if _, ok := a[0].(string); !ok {
			goto check_1
		}
		return p.GetMatchJson__SWIG_0(a[0].(string))
	}
check_1:
	if argc == 1 {
		return p.GetMatchJson__SWIG_1(a[0].(MapStringString))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrProvider) GetMatchForDeviceId(arg2 string) (_swig_ret Match) {
	var swig_r Match
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Match)(SwigcptrMatch(C._wrap_Provider_getMatchForDeviceId_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_37)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrProvider) FindProfiles__SWIG_0(arg2 string, arg3 string) (_swig_ret Profiles) {
	var swig_r Profiles
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Profiles)(SwigcptrProfiles(C._wrap_Provider_findProfiles__SWIG_0_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_38)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_39)(unsafe.Pointer(&_swig_i_2)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (arg1 SwigcptrProvider) FindProfiles__SWIG_1(arg2 string, arg3 string, arg4 Profiles) (_swig_ret Profiles) {
	var swig_r Profiles
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4.Swigcptr()
	swig_r = (Profiles)(SwigcptrProfiles(C._wrap_Provider_findProfiles__SWIG_1_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_40)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_41)(unsafe.Pointer(&_swig_i_2)), C.uintptr_t(_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (p SwigcptrProvider) FindProfiles(a ...interface{}) Profiles {
	argc := len(a)
	if argc == 2 {
		return p.FindProfiles__SWIG_0(a[0].(string), a[1].(string))
	}
	if argc == 3 {
		return p.FindProfiles__SWIG_1(a[0].(string), a[1].(string), a[2].(Profiles))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrProvider) ReloadFromFile() {
	_swig_i_0 := arg1
	C._wrap_Provider_reloadFromFile_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrProvider) ReloadFromMemory(arg2 string, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_Provider_reloadFromMemory_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0), *(*C.swig_type_42)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrProvider) GetCacheHits() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Provider_getCacheHits_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProvider) GetCacheMisses() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Provider_getCacheMisses_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProvider) GetCacheMaxIterations() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_Provider_getCacheMaxIterations_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrProvider) GetIsThreadSafe() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_Provider_getIsThreadSafe_FiftyOneDegreesPatternV3_c8b75028699b5788(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewProvider__SWIG_6(arg1 string, arg2 string, arg3 int, arg4 int, arg5 bool) (_swig_ret Provider) {
	var swig_r Provider
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	swig_r = (Provider)(SwigcptrProvider(C._wrap_new_Provider__SWIG_6_FiftyOneDegreesPatternV3_c8b75028699b5788(*(*C.swig_type_43)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_44)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C._Bool(_swig_i_4))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func NewProvider(a ...interface{}) Provider {
	argc := len(a)
	if argc == 1 {
		return NewProvider__SWIG_0(a[0].(string))
	}
	if argc == 2 {
		if _, ok := a[1].(string); !ok {
			goto check_2
		}
		return NewProvider__SWIG_1(a[0].(string), a[1].(string))
	}
check_2:
	if argc == 2 {
		return NewProvider__SWIG_3(a[0].(string), a[1].(VectorString))
	}
	if argc == 3 {
		return NewProvider__SWIG_5(a[0].(string), a[1].(int), a[2].(int))
	}
	if argc == 4 {
		if _, ok := a[1].(SwigcptrVectorString); !ok {
			goto check_5
		}
		return NewProvider__SWIG_4(a[0].(string), a[1].(VectorString), a[2].(int), a[3].(int))
	}
check_5:
	if argc == 4 {
		return NewProvider__SWIG_2(a[0].(string), a[1].(string), a[2].(int), a[3].(int))
	}
	if argc == 5 {
		return NewProvider__SWIG_6(a[0].(string), a[1].(string), a[2].(int), a[3].(int), a[4].(bool))
	}
	panic("No match for overloaded function call")
}

type Provider interface {
	Swigcptr() uintptr
	SwigIsProvider()
	GetHttpHeaders() (_swig_ret VectorString)
	GetAvailableProperties() (_swig_ret VectorString)
	GetDataSetName() (_swig_ret string)
	GetDataSetFormat() (_swig_ret string)
	GetDataSetPublishedDate() (_swig_ret string)
	GetDataSetNextUpdateDate() (_swig_ret string)
	GetDataSetSignatureCount() (_swig_ret int)
	GetDataSetDeviceCombinations() (_swig_ret int)
	GetMatch(a ...interface{}) Match
	GetMatchJson(a ...interface{}) string
	GetMatchForDeviceId(arg2 string) (_swig_ret Match)
	FindProfiles(a ...interface{}) Profiles
	ReloadFromFile()
	ReloadFromMemory(arg2 string, arg3 int)
	GetCacheHits() (_swig_ret int)
	GetCacheMisses() (_swig_ret int)
	GetCacheMaxIterations() (_swig_ret int)
	GetIsThreadSafe() (_swig_ret bool)
}


