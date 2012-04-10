package gobject

/*#include "go-gobject.h" */
// #cgo pkg-config: gobject-2.0
import "C"

import "unsafe"
import "reflect"
import "runtime"
import "time"
import "strings"

func GetTypeFromInstance(obj unsafe.Pointer) GType {
	return GType(C._get_type_from_instance(obj))
}

// Base interface which all gobject like structures must implement
// to be compatible with all gobject functions - public interfaces of gobject
type ObjectLike interface {
	ToNative() unsafe.Pointer
	Connect(s string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError)
	Set(map[string]interface{})
	Get([]string) map[string]interface{}
}

//-----------------------------------------------------------------------
// GObject interfaces
//-----------------------------------------------------------------------
// Convenient struct - singleton
type gobject struct {
	object unsafe.Pointer
}

func (self gobject) ToNative() unsafe.Pointer {
	return self.object
}

func (self gobject) Connect(name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	return Connect(self, name, f, data...)
}

func (self gobject) Set(properties map[string]interface{}) {
	Set(self, properties)
}

func (self gobject) Get(properties []string) map[string]interface{} {
	return Get(self, properties)
}

func New(typ GType, properties map[string]interface{}) ObjectLike {
	obj := gobject{}
	obj.object = C.new_GObject(C.GType(typ))

	Set(obj, properties)

	return obj

}

func FindProperty(obj ObjectLike, propName string) *GParamSpec {
	cpn := GString(propName)
	defer cpn.Free()

	oc := C.get_object_class(obj.ToNative())
	pspec := C.g_object_class_find_property(oc, (*C.gchar)(cpn.GetPtr()))
	if pspec == nil {
		return nil
	}
	gpspec := GParamSpec{pspec}
	return &gpspec
}

func Set(obj ObjectLike, properties map[string]interface{}) {
	if len(properties) == 0 {
		return
	}

	for p, v := range properties {
		SetProperty(obj, p, v)
	}
}

func Get(obj ObjectLike, properties []string) map[string]interface{} {
	m := map[string]interface{}{}
	for _, p := range properties {
		gv := GetProperty(obj, p)
		if gv == nil {
			continue
		}
		defer gv.Free()
		val, err := ConvertToGo(gv.GetPtr(), gv.gtype)
		if err == nil {
			m[p] = val
		}
	}
	return m
}

func SetProperty(obj ObjectLike, propName string, value interface{}) {
	cpn := GString(propName)
	defer cpn.Free()
	v := ConvertToC(value)
	defer v.Free()

	C.g_object_set_property(C.to_GObject(obj.ToNative()), (*C.gchar)(cpn.GetPtr()),
		(*C.GValue)(v.ToCGValue()))

}

func GetProperty(obj ObjectLike, propName string) *GValue {
	pspec := FindProperty(obj, propName)
	if pspec == nil {
		return nil
	}
	gv := CreateCGValue(pspec.GetType())
	pn := GString(propName)
	defer pn.Free()
	C.g_object_get_property(C.to_GObject(obj.ToNative()),
		(*C.gchar)(pn.GetPtr()),
		(*C.GValue)(gv.ToCGValue()))
	return gv
}

func Ref(obj ObjectLike) {
	C.g_object_ref(C.gpointer(obj.ToNative()))
}

func Unref(obj ObjectLike) {
	C.g_object_unref(C.gpointer(obj.ToNative()))
}

func SignalLookup(name string, objectType GType) uint32 {
	// If exists, remove detail part of the signal
	var canonicalName string = name
	if index := strings.Index(name, ":"); index != -1 {
		canonicalName = name[:index]
	}
	n := GString(canonicalName)
	defer n.Free()
	s := C.g_signal_lookup((*C.gchar)(n.GetPtr()), C.GType(objectType))
	return uint32(s)
}

func GetUniqueID() int64 {
	return time.Now().UnixNano()
}

type SignalError struct {
	Err string
}

func (self SignalError) Error() string {
	return self.Err
}

func createClosure(f interface{}, data ...interface{}) ClosureFunc {
	cfunc := reflect.ValueOf(f)

	// cfunc is not function, then there is nothing to do
	if cfunc.Kind() != reflect.Func {
		return nil
	}

	// Test number and type of return arguments
	// If return value is bool and is the only one, then
	// return that value to handler.
	// For anything else automaticaly return true
	var isTrue bool = true
	if numReturn := cfunc.Type().NumOut(); numReturn == 1 {
		rt := cfunc.Type().Out(0)
		if rt.Kind() == reflect.Bool {
			isTrue = false
		}
	}

	// Do we have variadic function for additional arguments
	var argslice bool = false
	if cfunc.Type().IsVariadic() {
		argslice = true
	}

	return func(additionalArgs []interface{}) bool {
		// Additional data
		// Create slice of reflect.Value data
		var args = make([]reflect.Value, 0)
		if len(data) > 0 {
			for _, d := range data {
				args = append(args, reflect.ValueOf(d))
			}
		}
		if argslice == true {
			for _, i := range additionalArgs {
				args = append(args, reflect.ValueOf(i))
			}
		}
		res := cfunc.Call(args)
		if isTrue {
			return isTrue
		}
		return res[0].Bool()
	}
}

func CreateCustomClosure(f interface{}, data ...interface{}) (ClosureFunc, int64) {
	cFunc := createClosure(f, data...)

	return cFunc, GetUniqueID()
}

func Connect(obj ObjectLike, name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	s_id := SignalLookup(name, GetTypeFromInstance(obj.ToNative()))

	if s_id == 0 {
		return nil, &SignalError{"Signal not found"}
	}
	uid := GetUniqueID()
	c := createClosure(f, data...)
	cloEl := RegisterHandler(obj, name, uid, c)
	return cloEl, nil
}

func IsObjectFloating(obj ObjectLike) bool {
	b := C.g_object_is_floating(C.gpointer(obj.ToNative()))
	return GoBool(unsafe.Pointer(&b))
}

func RefSink(obj ObjectLike) {
	C.g_object_ref_sink(C.gpointer(obj.ToNative()))
}

func ClearObject(obj ObjectLike) {
	C._g_clear_object(obj.ToNative())
}

// WeakRef closure stack
type WeakClosureFunc func()

var _weakClosures map[int64]WeakClosureFunc

func WeakRef(obj ObjectLike, f WeakClosureFunc) int64 {
	uid := GetUniqueID()
	_weakClosures[uid] = f
	o := C.to_GObject(obj.ToNative())
	C._g_object_weak_ref(o, C.gint64(uid))

	return uid
}

//export weak_ref_callback
func weak_ref_callback(data C.gpointer, obj *C.GObject) {
	cint := int64(*((*C.gint64)(data)))

	if f, ok := _weakClosures[cint]; ok {
		// Call WeakClosureFunc to clean up something
		f()

		// Remove this closure from _weakClosures
		delete(_weakClosures, cint)
	}

	// Free data
	C.destroy_data(C.gpointer(data))
}

// BoxedLike interface for Boxed values
type BoxedLike interface {
	GetBoxType() GType
}

// GBinding type
type GBinding struct {
	object *C.GBinding
}

// Clear GBinding struct when it goes out of reach
func gbindingFinalizer(gb *GBinding) {
	runtime.SetFinalizer(gb, func(gb *GBinding) { Unref(gb) })
}

// Conversion funcs
func gbindingFromNative(obj unsafe.Pointer) interface{} {
	gb := &GBinding{}
	gb.object = C.to_GBinding(obj)

	if IsObjectFloating(gb) {
		RefSink(gb)
	} else {
		Ref(gb)
	}
	gbindingFinalizer(gb)

	return gb
}

func (self GBinding) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self GBinding) Connect(name string, f interface{}, data ...interface{}) (*ClosureElement, *SignalError) {
	return Connect(self, name, f, data...)
}

func (self GBinding) Set(properties map[string]interface{}) {
	Set(self, properties)
}

func (self GBinding) Get(properties []string) map[string]interface{} {
	return Get(self, properties)
}

func BindProperty(source ObjectLike, sourceProperty string, dest ObjectLike, destProperty string, g_BindingFlags int) *GBinding {
	sp := GString(sourceProperty)
	defer sp.Free()
	dp := GString(destProperty)
	defer dp.Free()
	o := C.g_object_bind_property(C.gpointer(source.ToNative()), (*C.gchar)(sp.GetPtr()),
		C.gpointer(dest.ToNative()), (*C.gchar)(dp.GetPtr()), C.GBindingFlags(g_BindingFlags))

	return gbindingFromNative(unsafe.Pointer(o)).(*GBinding)
}

var GBindingFlags gbindingFlags

type gbindingFlags struct {
	DEFAULT        int
	BIDIRECTIONAL  int
	SYNC_CREATE    int
	INVERT_BOOLEAN int
}

// Foundamental types {{{

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

// Definition of GParamSpec Type
type GParamSpec struct {
	object *C.GParamSpec
}

// Conversion funcs
func newGParamSpecFromNative(obj unsafe.Pointer) interface{} {
	p := &GParamSpec{}
	p.object = C.to_GParamSpec(obj)

	return p
}

func nativeFromGParamSpec(ps interface{}) *GValue {
	if p, ok := ps.(*GParamSpec); ok {
		gv := CreateCGValue(G_TYPE_PARAM, p.ToNative())
		return gv
	}
	return nil
}

func (self GParamSpec) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self GParamSpec) ValueSetDefaults(val *GValue) {
	C.g_param_value_set_default((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
}

func (self GParamSpec) ValueDefaults(val *GValue) bool {
	v := C.g_param_value_defaults((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
	return GoBool(unsafe.Pointer(&v))
}

func (self GParamSpec) ValueValidate(val *GValue) bool {
	v := C.g_param_value_validate((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
	return GoBool(unsafe.Pointer(&v))
}

func (self GParamSpec) ValueConvert(src, dest *GValue, strict_validation bool) bool {
	sv := GBool(strict_validation)
	res := C.g_param_value_convert((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(src.ToCGValue()),
		(*C.GValue)(dest.ToCGValue()),
		*((*C.gboolean)(sv.GetPtr())))
	return GoBool(unsafe.Pointer(&res))
}

func (self GParamSpec) GetType() GType {
	return GType(C.get_type(self.object))
}

// End Foundamental types }}}

// Closures support {{{

type ClosureFunc func(args []interface{}) bool

// This gets returned by RegisterHandler
type ClosureElement struct {
	id         int64
	handler_id uint32
}

var _closures map[int64]ClosureFunc

// RegisterHandler registers closure and returns ClosureElement with info for removal
func RegisterHandler(obj ObjectLike, name string, id int64, f ClosureFunc) *ClosureElement {
	// Register handler to our gobject system
	_closures[id] = f

	// Register handler in gobject system
	s := GString(name)
	defer s.Free()
	h_id := C.connect_to_signal(obj.ToNative(),
		(*C.gchar)(s.GetPtr()), C.gint64(id))

	return &ClosureElement{id, uint32(h_id)}
}

// RemoveHandler removes closure from our map and from gobject signal system
func RemoveHandler(obj ObjectLike, cel *ClosureElement) {
	if _, ok := _closures[cel.id]; ok {
		delete(_closures, cel.id)
		C.g_signal_handler_disconnect(C.gpointer(obj.ToNative()), C.gulong(cel.handler_id))
	}

}
// End Closures support }}}

// GType support {{{

type GType int64

type GValueError struct {
	Err string
}

func (e GValueError) Error() string {
	return e.Err
}

// Function interface for converting C values to Go
type FuncToGo func(unsafe.Pointer) interface{}

// GoTypes registration map
// This map holds gobject types as GType
// and conversion function to Go type
// Every module has to register it's own type and conversion function
var gtypes map[GType]FuncToGo

// IsObjectType returns true if given type is derived from GObject
func IsObjectType(t GType) bool {
	b := C.is_type_object(C.GType(t))
	return GoBool(unsafe.Pointer(&b))
}

// IsEnumType returns true if given type is derived from GEnum
func IsEnumType(t GType) bool {
	b := C._is_type_enum(C.GType(t))
	return GoBool(unsafe.Pointer(&b))
}

// IsFlagsType returns true if given type is derived from GFlags
func IsFlagsType(t GType) bool {
	b := C._is_type_flags(C.GType(t))
	return GoBool(unsafe.Pointer(&b))
}

// IsBoxedType returns true if given type is derived from GBoxed
func IsBoxedType(t GType) bool {
	b := C.is_type_boxed(C.GType(t))
	return GoBool(unsafe.Pointer(&b))
}

// Get Type ID
func GetTypeID(obj unsafe.Pointer) GType {
	t := C.getTypeId(obj)
	return GType(t)
}

// Get Type Name
func GetTypeName(obj unsafe.Pointer) string {
	tn := C.getTypeName(obj)
	return GoString(unsafe.Pointer(tn))
}

func RegisterCType(typename GType, f FuncToGo) {
	if gtypes == nil {
		gtypes = make(map[GType]FuncToGo)
	}
	gtypes[typename] = f
}

func ConvertToGo(obj unsafe.Pointer, typeid ...GType) (interface{}, error) {
	if obj == nil {
		return nil, GValueError{"Conversion Error"}
	}
	var id GType

	if len(typeid) > 0 {
		id = typeid[0]
	} else {
		id = GetTypeID(obj)
	}
	if f, ok := gtypes[id]; ok {
		res := f(obj)
		return res, nil
	}

	if IsFlagsType(id) {
		return uint(*((*C.guint)(obj))), nil
	}

	if IsEnumType(id) {
		e := *((*C.gint)(obj))
		return int(e), nil
	}

	return nil, GValueError{"Unknown Type"}
}

// Function interface for converting Go values to C
type FuncToC func(interface{}) *GValue

var ctypes map[GType]FuncToC

func RegisterGoType(typeid GType, f FuncToC) {
	if ctypes == nil {
		ctypes = make(map[GType]FuncToC)
	}
	ctypes[typeid] = f
}

// GValue
type GValue struct {
	gtype GType
	value *C.GValue
}

func NewGValueFromNative(gv unsafe.Pointer) *GValue {
	cgv := (*C.GValue)(gv)
	newGV := &GValue{0, cgv}
	newGV.ReInitializeType()
	return newGV
}

func (self GValue) ToCGValue() unsafe.Pointer {
	return unsafe.Pointer(self.value)
}

func (self *GValue) ReInitializeType() {
	t := C.getGValueType(self.value)
	self.gtype = GType(t)
}

func (self GValue) GetTypeName() string {
	tn := C.getGValueTypeName(self.value)
	return GoString(unsafe.Pointer(tn))
}

func (self GValue) GetTypeID() GType {
	return self.gtype
}

func (self GValue) GetPtr() unsafe.Pointer {
	t := self.gtype
	switch {
	case IsObjectType(t):
		return unsafe.Pointer(C.g_value_get_object(self.value))
	case IsEnumType(t):
		e := C.g_value_get_enum(self.value)
		return unsafe.Pointer(&e)
	case IsFlagsType(t):
		f := C.g_value_get_flags(self.value)
		return unsafe.Pointer(&f)
	case IsBoxedType(t):
		b := C.g_value_get_boxed(self.value)
		return unsafe.Pointer(b)
	case t == G_TYPE_STRING:
		return unsafe.Pointer(C.g_value_get_string(self.value))
	case t == G_TYPE_BOOLEAN:
		b := C.g_value_get_boolean(self.value)
		return unsafe.Pointer(&b)
	case t == G_TYPE_CHAR:
		c := C.g_value_get_schar(self.value)
		return unsafe.Pointer(&c)
	case t == G_TYPE_INT:
		i := C.g_value_get_int(self.value)
		return unsafe.Pointer(&i)
	case t == G_TYPE_LONG:
		l := C.g_value_get_long(self.value)
		return unsafe.Pointer(&l)
	case t == G_TYPE_INT64:
		i := C.g_value_get_int64(self.value)
		return unsafe.Pointer(&i)
	case t == G_TYPE_UCHAR:
		c := C.g_value_get_uchar(self.value)
		return unsafe.Pointer(&c)
	case t == G_TYPE_UINT:
		i := C.g_value_get_uint(self.value)
		return unsafe.Pointer(&i)
	case t == G_TYPE_ULONG:
		l := C.g_value_get_ulong(self.value)
		return unsafe.Pointer(&l)
	case t == G_TYPE_UINT64:
		i := C.g_value_get_uint64(self.value)
		return unsafe.Pointer(&i)
	case t == G_TYPE_FLOAT:
		f := C.g_value_get_float(self.value)
		return unsafe.Pointer(&f)
	case t == G_TYPE_DOUBLE:
		d := C.g_value_get_double(self.value)
		return unsafe.Pointer(&d)
	case t == G_TYPE_PARAM:
		p := C.g_value_get_param(self.value)
		return unsafe.Pointer(p)
	}

	p := C.g_value_get_pointer(self.value)
	return unsafe.Pointer(p)
}

func (self GValue) Free() {
	C.g_value_unset(self.value)
}

func CreateCGValue(tn GType, object ...unsafe.Pointer) *GValue {
	var cv C.GValue
	C.g_value_init(&cv, C.GType(tn))

	// If no data, then return Gvalue initialized with default
	if len(object) == 0 {
		gv := GValue{tn, &cv}
		return &gv
	}
	obj := object[0]

	// Foundamental types are special
	switch {
	case IsObjectType(tn):
		C.g_value_set_object(&cv, C.gpointer(obj))
	case IsEnumType(tn):
		C.g_value_set_enum(&cv, *((*C.gint)(obj)))
	case IsFlagsType(tn):
		C.g_value_set_flags(&cv, *((*C.guint)(obj)))
	case IsBoxedType(tn):
		C.g_value_take_boxed(&cv, C.gconstpointer(obj))
	case tn == G_TYPE_STRING:
		C.g_value_take_string(&cv, (*C.gchar)(obj))
	case tn == G_TYPE_BOOLEAN:
		C.g_value_set_boolean(&cv, *((*C.gboolean)(obj)))
	case tn == G_TYPE_CHAR:
		C.g_value_set_schar(&cv, *((*C.gint8)(obj)))
	case tn == G_TYPE_INT:
		C.g_value_set_int(&cv, *((*C.gint)(obj)))
	case tn == G_TYPE_LONG:
		C.g_value_set_long(&cv, *((*C.glong)(obj)))
	case tn == G_TYPE_INT64:
		C.g_value_set_int64(&cv, *((*C.gint64)(obj)))
	case tn == G_TYPE_UCHAR:
		C.g_value_set_uchar(&cv, *((*C.guchar)(obj)))
	case tn == G_TYPE_UINT:
		C.g_value_set_uint(&cv, *((*C.guint)(obj)))
	case tn == G_TYPE_ULONG:
		C.g_value_set_ulong(&cv, *((*C.gulong)(obj)))
	case tn == G_TYPE_UINT64:
		C.g_value_set_uint64(&cv, *((*C.guint64)(obj)))
	case tn == G_TYPE_FLOAT:
		C.g_value_set_float(&cv, *((*C.gfloat)(obj)))
	case tn == G_TYPE_DOUBLE:
		C.g_value_set_double(&cv, *((*C.gdouble)(obj)))
	case tn == G_TYPE_PARAM:
		C.g_value_set_param(&cv, (*C.GParamSpec)(obj))
	case tn == G_TYPE_POINTER:
		C.g_value_set_pointer(&cv, C.gpointer(obj))
	}

	gv := GValue{tn, &cv}
	return &gv
}

func ConvertToC(gotype interface{}) *GValue {
	// Test first for fundamental types
	switch gotype.(type) {
	case string:
		return ctypes[G_TYPE_STRING](gotype.(string))
	case bool:
		return ctypes[G_TYPE_BOOLEAN](gotype.(bool))
	case int8:
		return ctypes[G_TYPE_CHAR](gotype.(byte))
	case int:
		return ctypes[G_TYPE_INT](gotype.(int))
	case int32:
		return ctypes[G_TYPE_LONG](gotype.(int32))
	case int64:
		return ctypes[G_TYPE_INT64](gotype.(int64))
	case byte:
		return ctypes[G_TYPE_UCHAR](gotype.(uint8))
	case uint:
		return ctypes[G_TYPE_UINT](gotype.(uint))
	case uint32:
		return ctypes[G_TYPE_ULONG](gotype.(uint32))
	case uint64:
		return ctypes[G_TYPE_UINT64](gotype.(uint64))
	case float32:
		return ctypes[G_TYPE_FLOAT](gotype.(float32))
	case float64:
		return ctypes[G_TYPE_DOUBLE](gotype.(float64))
	case *GParamSpec:
		return ctypes[G_TYPE_PARAM](gotype.(*GParamSpec))
	case ObjectLike:
		o := gotype.(ObjectLike)
		t := GetTypeFromInstance(o.ToNative())
		if f, ok := ctypes[t]; ok {
			return f(o)
		}
	case BoxedLike:
		o := gotype.(BoxedLike)
		t := o.GetBoxType()
		if f, ok := ctypes[t]; ok {
			return f(o)
		}
	}
	return nil
}

// End GType support }}}

func init() {
	C.g_type_init()
	_closures = make(map[int64]ClosureFunc)
	_weakClosures = make(map[int64]WeakClosureFunc)

    gtypes = make(map[GType]FuncToGo)
    ctypes = make(map[GType]FuncToC)

	// Foundamental type registration {{{

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

	// GParamSpec
	RegisterCType(G_TYPE_PARAM, newGParamSpecFromNative)
	RegisterGoType(G_TYPE_PARAM, nativeFromGParamSpec)

    // End foundamental type registration }}}


	// Initialize GBindingFlags
	GBindingFlags.DEFAULT = 0
	GBindingFlags.BIDIRECTIONAL = 1 << 0
	GBindingFlags.SYNC_CREATE = 1 << 1
	GBindingFlags.INVERT_BOOLEAN = 1 << 2
}

// EXPORTED FUNCS {{{

// SimpleGoMarshal
// Our true callback which gets called
// for all connected signals.
// It call appropriate closure which it finds in closures map
//export simple_go_marshal
func simple_go_marshal(closure *C.GClosure,
	returnValue *C.GValue,
	n_param_values C.guint,
	paramValues C.constgvalue,
	invocationHint C.gpointer,
	marshalData C.gpointer) {

//	c := (*C.GClosure)(closure)
	id := int64(*((*C.gint64)(closure.data)))

	argslice := make([]interface{}, int(n_param_values))
	array := C.g_array_from_GValues(paramValues, n_param_values)
	for i := 0; i < int(n_param_values); i++ {
		v := C.get_index(array, C.guint(i))
		gv := GValue{0, &v}
		gv.ReInitializeType()
		t, e := ConvertToGo(gv.GetPtr(), gv.GetTypeID())
		if e == nil {
			argslice[i] = t
		} else {
			argslice[i] = nil
		}
	}

	if h, ok := _closures[id]; ok {
		h(argslice)
	}

	C.free_array(array)
}
// EXPORTED FUNCS }}}
