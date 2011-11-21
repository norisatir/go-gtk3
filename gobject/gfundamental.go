package gobject

/*
#include <glib-object.h>

*/
import "C"

import "unsafe"

// Basic gobject types
const (
	G_TYPE_NONE      GType = 1 << 2
	G_TYPE_INTERFACE GType = 2 << 2
	G_TYPE_CHAR      GType = 3 << 2
	G_TYPE_UCHAR     GType = 4 << 2
	G_TYPE_BOOLEAN   GType = 5 << 2
	G_TYPE_INT       GType = 6 << 2
	G_TYPE_UINT      GType = 7 << 2
	G_TYPE_LONG      GType = 8 << 2
	G_TYPE_ULONG     GType = 9 << 2
	G_TYPE_INT64     GType = 10 << 2
	G_TYPE_UINT64    GType = 11 << 2
	G_TYPE_ENUM      GType = 12 << 2
	G_TYPE_FLAGS     GType = 13 << 2
	G_TYPE_FLOAT     GType = 14 << 2
	G_TYPE_DOUBLE    GType = 15 << 2
	G_TYPE_STRING    GType = 16 << 2
	G_TYPE_POINTER   GType = 17 << 2
	G_TYPE_BOXED     GType = 18 << 2
	G_TYPE_PARAM     GType = 19 << 2
	G_TYPE_OBJECT    GType = 20 << 2
	G_TYPE_VARIANT   GType = 21 << 2
)

func init() {
	// Foundamental register

	// String
	RegisterGoType(G_TYPE_STRING, GString)
	RegisterCType(G_TYPE_STRING, to_GoString)

	// Bool
	RegisterGoType(G_TYPE_BOOLEAN, GBool)
	RegisterCType(G_TYPE_BOOLEAN, to_GoBool)

	// Char
	RegisterGoType(G_TYPE_CHAR, GChar)
	RegisterCType(G_TYPE_CHAR, to_GoChar)

	// Int
	RegisterGoType(G_TYPE_INT, GInt)
	RegisterCType(G_TYPE_INT, to_GoInt)

	// Int32
	RegisterGoType(G_TYPE_LONG, GLong)
	RegisterCType(G_TYPE_LONG, to_GoLong)

	// Int64
	RegisterGoType(G_TYPE_INT64, GInt64)
	RegisterCType(G_TYPE_INT64, to_GoInt64)

	// UChar
	RegisterGoType(G_TYPE_UCHAR, GUChar)
	RegisterCType(G_TYPE_UCHAR, to_GoUChar)

	// UInt
	RegisterGoType(G_TYPE_UINT, GUInt)
	RegisterCType(G_TYPE_UINT, to_GoUInt)

	// UInt32
	RegisterGoType(G_TYPE_ULONG, GULong)
	RegisterCType(G_TYPE_ULONG, to_GoUlong)

	// Uint64
	RegisterGoType(G_TYPE_UINT64, GUInt64)
	RegisterCType(G_TYPE_UINT64, to_GoUInt64)

	// Float
	RegisterGoType(G_TYPE_FLOAT, GFloat)
	RegisterCType(G_TYPE_FLOAT, to_GoFloat)

	// Double
	RegisterGoType(G_TYPE_DOUBLE, GDouble)
	RegisterCType(G_TYPE_DOUBLE, to_GoDouble)

	// Object
	RegisterGoType(G_TYPE_OBJECT, GObject)
	RegisterCType(G_TYPE_OBJECT, to_GoObject)
}

// Foundamental types
func GString(val interface{}) *GValue {
	s := val.(string)
	cs := C.CString(s)
	return CreateCGValue(G_TYPE_STRING, unsafe.Pointer((*C.gchar)(cs)))
}

func GoString(str unsafe.Pointer) string {
	return C.GoString((*C.char)(str))
}

func to_GoString(str unsafe.Pointer) interface{} {
	return GoString(str)
}

func GBool(val interface{}) *GValue {
	b := val.(bool)
	var cb C.gboolean
	if b {
		cb = C.gboolean(1)
	} else {
		cb = C.gboolean(0)
	}
	return CreateCGValue(G_TYPE_BOOLEAN, unsafe.Pointer(&cb))
}

func GoBool(b unsafe.Pointer) bool {
	gb := *((*C.gboolean)(b))
	if gb > 0 {
		return true
	}
	return false
}

func to_GoBool(b unsafe.Pointer) interface{} {
	return GoBool(b)
}

func GChar(val interface{}) *GValue {
	b := val.(int8)
	by := C.gchar(b)
	return CreateCGValue(G_TYPE_CHAR, unsafe.Pointer(&by))
}

func GoChar(chr unsafe.Pointer) int8 {
	return int8(*((*C.gchar)(chr)))
}

func to_GoChar(chr unsafe.Pointer) interface{} {
	return GoChar(chr)
}

func GInt(val interface{}) *GValue {
	i := val.(int)
	iv := C.gint(i)
	return CreateCGValue(G_TYPE_INT, unsafe.Pointer(&iv))
}

func GoInt(gint unsafe.Pointer) int {
	return int(*((*C.gint)(gint)))
}

func to_GoInt(gint unsafe.Pointer) interface{} {
	return GoInt(gint)
}

func GLong(val interface{}) *GValue {
	i := val.(int32)
	iv := C.glong(i)
	return CreateCGValue(G_TYPE_LONG, unsafe.Pointer(&iv))
}

func GoLong(glong unsafe.Pointer) int32 {
	return int32(*((*C.glong)(glong)))
}

func to_GoLong(glong unsafe.Pointer) interface{} {
	return GoLong(glong)
}

func GInt64(val interface{}) *GValue {
	i := val.(int64)
	iv := C.gint64(i)
	return CreateCGValue(G_TYPE_INT64, unsafe.Pointer(&iv))
}

func GoInt64(gint64 unsafe.Pointer) int64 {
	return int64(*((*C.gint64)(gint64)))
}

func to_GoInt64(gint64 unsafe.Pointer) interface{} {
	return GoInt64(gint64)
}

func GUChar(val interface{}) *GValue {
	uc := val.(byte)
	uv := C.guchar(uc)
	return CreateCGValue(G_TYPE_UCHAR, unsafe.Pointer(&uv))
}

func GoUChar(guchar unsafe.Pointer) byte {
	return byte(*((*C.guchar)(guchar)))
}

func to_GoUChar(guchar unsafe.Pointer) interface{} {
	return GoUChar(guchar)
}

func GUInt(val interface{}) *GValue {
	ui := val.(uint)
	uv := C.guint(ui)
	return CreateCGValue(G_TYPE_UINT, unsafe.Pointer(&uv))
}

func GoUInt(guint unsafe.Pointer) uint {
	return uint(*((*C.guint)(guint)))
}

func to_GoUInt(guint unsafe.Pointer) interface{} {
	return GoUInt(guint)
}

func GULong(val interface{}) *GValue {
	ui := val.(uint32)
	ul := C.gulong(ui)
	return CreateCGValue(G_TYPE_ULONG, unsafe.Pointer(&ul))
}

func GoULong(gulong unsafe.Pointer) uint32 {
	return uint32(*((*C.gulong)(gulong)))
}

func to_GoUlong(gulong unsafe.Pointer) interface{} {
	return GoULong(gulong)
}

func GUInt64(val interface{}) *GValue {
	ui := val.(uint64)
	ul := C.guint64(ui)
	return CreateCGValue(G_TYPE_UINT64, unsafe.Pointer(&ul))
}

func GoUInt64(guint64 unsafe.Pointer) uint64 {
	return uint64(*((*C.guint64)(guint64)))
}

func to_GoUInt64(guint64 unsafe.Pointer) interface{} {
	return GoUInt64(guint64)
}

func GFloat(val interface{}) *GValue {
	fl := val.(float32)
	gf := C.gfloat(fl)
	return CreateCGValue(G_TYPE_FLOAT, unsafe.Pointer(&gf))
}

func GoFloat(gfloat unsafe.Pointer) float32 {
	return float32(*((*C.gfloat)(gfloat)))
}

func to_GoFloat(gfloat unsafe.Pointer) interface{} {
	return GoFloat(gfloat)
}

func GDouble(val interface{}) *GValue {
	fd := val.(float64)
	gd := C.gdouble(fd)
	return CreateCGValue(G_TYPE_DOUBLE, unsafe.Pointer(&gd))
}

func GoDouble(gdouble unsafe.Pointer) float64 {
	return float64(*((*C.gdouble)(gdouble)))
}

func to_GoDouble(gdouble unsafe.Pointer) interface{} {
	return GoDouble(gdouble)
}

func GObject(val interface{}) *GValue {
	o := val.(ObjectLike)
	return CreateCGValue(G_TYPE_OBJECT, o.ToNative())
}

func GoObject(obj unsafe.Pointer) ObjectLike {
	return gobject{obj}
}

func to_GoObject(obj unsafe.Pointer) interface{} {
	return GoObject(obj)
}
